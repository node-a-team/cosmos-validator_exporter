package commit

import (
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
)

func CommitStatus(blockHeight int) t.CommitStatus {

	var commitStatus t.CommitStatus

	var blockHeightInt64 int64 = int64(blockHeight)

	info, _ := t.Client.Commit(&blockHeightInt64)
	infoMarshal, _ := json.Marshal(info)
	json.Unmarshal(infoMarshal, &commitStatus)

	return commitStatus

}

func ValidatorPrecommitStatus(commitStatus t.CommitStatus, consHexAddr string) int {

	// validatorCommitStatus(0: false, 1 : true)
	validatorCommitStatus := 0

	precommitData := commitStatus.Signed_header.Commit.Precommits
//	precommitData := commitStatus.Result.Signed_header.Commit.Precommits

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

	precommitData := commitStatus.Signed_header.Commit.Precommits
//	precommitData := commitStatus.Result.Signed_header.Commit.Precommits
	totalCount, precommitCount := len(precommitData), len(precommitData)

	for _, value := range precommitData {

		if value.Validator_address == "" {
			precommitCount--
		}
	}

	precommitRate = float64(precommitCount) / float64(totalCount)

	return precommitRate
}
