package validators

import (
	utils "github.com/node-a-team/cosmos-validator_exporter/function/utils"
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
	"os/exec"
)

var ()

func ValidatorDelegatorNumber(operatorAddr string, accountAddr string) (delegatorCount float64, selfDelegationAmount float64) {

	var validatorDelegationStatus []t.ValidatorDelegationStatus

	delegatorCount, selfDelegationAmount = 0.0, 0.0

	cmd := "curl -s -XGET " + t.RestServer + "/staking/validators/" + operatorAddr + "/delegations" + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &validatorDelegationStatus)

	delegatorCount = float64(len(validatorDelegationStatus))

	for _, value := range validatorDelegationStatus {
		if value.Delegator_Address == accountAddr {
			selfDelegationAmount = utils.StringToFloat64(value.Shares)
		}
	}

	return delegatorCount, selfDelegationAmount
}
