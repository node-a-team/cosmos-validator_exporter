package validators

import (
	"encoding/json"
	"fmt"
	t "github.com/node-a-team/cosmos-validator_exporter/types"
	"os/exec"
)

var ()

func ValidatorsStatus() map[string][]string {

	var validatorsStatus []t.ValidatorStatus
	var validators map[string][]string = make(map[string][]string)

	cmd := "curl -s -XGET " + t.RestServer + "/staking/validators?status=bonded&page=1&limit=1" + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &validatorsStatus)

	for _, value := range validatorsStatus {
		validators[value.Consensus_pubkey] = []string{value.Description.Moniker, value.Operator_address, fmt.Sprint(value.Jailed), value.Tokens, value.Delegator_shares, value.Commission.Rate, value.Commission.Max_rate, value.Commission.Max_change_rate, value.Commission.Update_time, value.Min_self_delegation, value.Unbonding_height, value.Unbonding_time, value.Description.Identity, value.Description.Website, value.Description.Details}

	}

	return validators
}
