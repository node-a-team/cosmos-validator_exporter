package staking

import (
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
	"os/exec"
	"strconv"
)

var ()

func GetStakingPool() (notBondedTokens float64, bondedTokens float64) {

	var stakingPoolStatus t.StakingPool

	cmd := "curl -s " + t.RestServer + "/staking/pool"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &stakingPoolStatus)

	notBondedTokens, _ = strconv.ParseFloat(stakingPoolStatus.Not_bonded_tokens, 64)
	bondedTokens, _ = strconv.ParseFloat(stakingPoolStatus.Bonded_tokens, 64)

	return notBondedTokens, bondedTokens

}
