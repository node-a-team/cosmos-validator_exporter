package validators

import (
	"encoding/json"
	"fmt"
	t "github.com/node-a-team/cosmos-validator_exporter/types"
	"os/exec"
)

var ()

func ValidatorsStatus() map[string][]string {

	var validatorsStatus t.ValidatorStatus
	var validators map[string][]string = make(map[string][]string)

	cmd := "curl -s -XGET " + t.RestServer + "/staking/validators?status=bonded" + " -H \"accept:application/json\""

	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &validatorsStatus)


	for _, value := range validatorsStatus.Result {
		validators[value.Consensus_pubkey] = []string{value.Description.Moniker, value.Operator_address, fmt.Sprint(value.Jailed), value.Tokens, value.Delegator_shares, value.Commission.Commission_rates.Rate, value.Commission.Commission_rates.Max_rate, value.Commission.Commission_rates.Max_change_rate, value.Commission.Update_time, value.Min_self_delegation, value.Unbonding_height, value.Unbonding_time, value.Description.Identity, value.Description.Website, value.Description.Details}
	}

	return validators
}
