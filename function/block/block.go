package block

import (
	keyutil "github.com/node-a-team/cosmos-validator_exporter/function/keyutil"
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

var ()

func BlockStatus() t.BlockStatus {

	var blockStatus t.BlockStatus

	cmd := "curl -s " + t.RpcServer + "/block"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &blockStatus)

	currentBlockHeight, _ := strconv.Atoi(blockStatus.Result.Block.Header.Height)

	// 현재 precommit 정보와 현재 blockHeight를 맞추기 위해 이전 블록 정보로 표시
	blockStatus = previousBlockStatus(currentBlockHeight - 1)

	return blockStatus

}

func previousBlockStatus(blockHeight int) t.BlockStatus {

	var blockStatus t.BlockStatus

	cmd := "curl -s " + t.RpcServer + "/block?height=" + fmt.Sprint(blockHeight)
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &blockStatus)

	return blockStatus

}

func CalcBlockTime(currentBlockStatus t.BlockStatus) float64 {

	var blockTime float64

	currentBlockHeight, _ := strconv.Atoi(currentBlockStatus.Result.Block.Header.Height)
	previousBlockHeight := currentBlockHeight - 1

	currentBlockTime := currentBlockStatus.Result.Block.Header.Time
	previousBlockTime := previousBlockStatus(previousBlockHeight).Result.Block.Header.Time

	if previousBlockTime.IsZero() {
		blockTime = 0.0
	} else {
		blockTime = float64(currentBlockTime.Sub(previousBlockTime)) / 1000000000
		blockTime, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", blockTime), 64)
	}

	return blockTime
}

func ProposerMoniker(propserAddress string, validatorsetsStatus map[string][]string, validatorsStatus map[string][]string) string {

	var proposerMoniker string
	keys := keyutil.RunFromHex(propserAddress)

	for validator_pubKey := range validatorsetsStatus {
		if keys[4] == validatorsetsStatus[validator_pubKey][0] {
			proposerMoniker = validatorsStatus[validator_pubKey][0]
		}
	}

	return proposerMoniker

}
