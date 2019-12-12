package validators

import (
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	keyutil "github.com/node-a-team/cosmos-validator_exporter/function/keyutil"
	utils "github.com/node-a-team/cosmos-validator_exporter/function/utils"

	"fmt"
	"log"

	"encoding/json"
	"os/exec"
	"strings"

	"bufio"
	"encoding/csv"
	"os"
)

var ()

func ValidatorAccount(accountAddr string) (balances []t.Coin, accountNumber float64) {

	cmd := "curl -s -XGET " + t.RestServer + "/auth/accounts/" + accountAddr + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()

	if strings.Index(string(out), "BaseVestingAccount") == -1 {

		// defaultAccount
		var accountStatus t.AccountStatus
		json.Unmarshal(out, &accountStatus)

		balances = accountStatus.Result.Value.Coins
		accountNumber = utils.StringToFloat64(accountStatus.Result.Value.Account_number)

	} else {

		// baseVestingAccount
		var baseVestingAccountStatus t.BaseVestingAccountStatus
		json.Unmarshal(out, &baseVestingAccountStatus)

		balances = baseVestingAccountStatus.Result.Value.BaseVestingAccount.BaseAccount.Coins
		accountNumber = utils.StringToFloat64(baseVestingAccountStatus.Result.Value.BaseVestingAccount.BaseAccount.Account_number)
	}

	return balances, accountNumber
}

func ValidatorsAccountNumber(blockHeight float64, validatorsetsStatus map[string][]string, validatorsStatus map[string][]string) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// csv
	csvFile, _ := os.OpenFile(homeDir+"/validatorsWalletAccounNumber.csv", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644) // 누적
	csvWriter := csv.NewWriter(bufio.NewWriter(csvFile))
	defer csvWriter.Flush()

	csvWriter.Write([]string{"> BlockHeight : " + fmt.Sprint(int(blockHeight))})
	csvWriter.Write([]string{"[ Moniker ]", "[ WalletAccount] ", "[ AccountAddress ]"})

	for validatorPubKey := range validatorsetsStatus {

		operatorAddress := validatorsStatus[validatorPubKey][1]
		accountAddress := keyutil.OperAddrToOtherAddr(operatorAddress)[0]

		moniker := validatorsStatus[validatorPubKey][0]
		_, validatorWalletAccountNumber := ValidatorAccount(accountAddress)

		csvWriter.Write([]string{moniker, fmt.Sprint(validatorWalletAccountNumber), accountAddress})
	}
}
