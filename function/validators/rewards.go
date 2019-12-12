package validators

import (
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
	"os/exec"
)

var ()

func ValidatorRewards(operatorAddr string) (rewards []t.Coin, commission []t.Coin) {

	var validatorRewards t.ValidatorRewards

	cmd := "curl -s -XGET " + t.RestServer + "/distribution/validators/" + operatorAddr + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &validatorRewards)


	rewards = validatorRewards.Result.Self_Bond_Rewards
	commission = validatorRewards.Result.Val_commission

	return rewards, commission
}
