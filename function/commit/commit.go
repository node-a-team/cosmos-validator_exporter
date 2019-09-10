package commit

import (
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
	"fmt"
	"os/exec"
)

func CommitStatus(blockHeight int) t.CommitStatus {

	var commitStatus t.CommitStatus

	cmd := "curl -s -XGET " + t.RpcServer + "/commit?height=" + fmt.Sprint(blockHeight) + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &commitStatus)

	return commitStatus

}

func ValidatorPrecommitStatus(commitStatus t.CommitStatus, consHexAddr string) int {

	// validatorCommitStatus(0: false, 1 : true)
	validatorCommitStatus := 0

	precommitData := commitStatus.Result.Signed_header.Commit.Precommits

	for _, value := range precommitData {

		if value.Validator_address == consHexAddr {
			validatorCommitStatus = 1
		}
	}

	return validatorCommitStatus
}

func PrecommitRate(commitStatus t.CommitStatus) float64 {

	// validatorCommitStatus(0: false, 1 : true)
	precommitRate := 0.0

	precommitData := commitStatus.Result.Signed_header.Commit.Precommits
	totalCount, precommitCount := len(precommitData), len(precommitData)

	for _, value := range precommitData {

		if value.Validator_address == "" {
			precommitCount--
		}
	}

	precommitRate = float64(precommitCount) / float64(totalCount)

	return precommitRate
}
