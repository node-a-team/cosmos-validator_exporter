package account

import (
        t "github.com/node-a-team/cosmos-validator_exporter/types"

	"os/exec"
        "encoding/json"
	"strconv"
)


var (
)


func Account(accaddr string) ([]string, []int64) {

	var coins []t.Coin
	var resultDenom []string
	var resultAmount []int64

	cmd := "curl -s -XGET " +t.RestServer +"/bank/balances/" +accaddr  +" -H \"accept:application/json\""
	fmt.Println(cmd)
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
        json.Unmarshal(out, &coins)

	resultDenom = make([]string, len(coins))
        resultAmount = make([]int64, len(coins))

	for i, value := range coins {
		resultDenom[i] = value.Denom
		resultAmount[i], _ = strconv.ParseInt(value.Amount, 10, 32)
	}

	return  resultDenom, resultAmount

}
