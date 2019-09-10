package validators

import (
	"encoding/json"
	t "github.com/node-a-team/cosmos-validator_exporter/types"
	"os/exec"
	"sort"
	"strconv"
)

var (
	validatorsetsStatus t.ValidatorsetsStatus

	// proposer[0] : priority 값
	// proposer[1] : priority 순위
	validatorsets map[string][]string = make(map[string][]string)
)

func ValidatorsetsStatus() (int, map[string][]string) {

	cmd := "curl -s -XGET " + t.RestServer + "/validatorsets/latest -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &validatorsetsStatus)

	for _, value := range validatorsetsStatus.Validators {
		validatorsets[value.Pub_key] = []string{value.Address, value.Voting_power, value.Proposer_priority, "0"}
	}

	validatorCount := len(validatorsets)
	validatorsetsSort := Sort(validatorsets)

	return validatorCount, validatorsetsSort
}

func Sort(mapValue map[string][]string) map[string][]string {

	keys := []string{}
	newMapValue := mapValue

	// key(moniker)들을 keys 배열에 추가
	for key := range mapValue {
		keys = append(keys, key)
	}

	// keys를 totalStake 기준으로 정렬
	sort.Slice(keys, func(i, j int) bool {
		a, _ := strconv.Atoi(mapValue[keys[i]][2])
		b, _ := strconv.Atoi(mapValue[keys[j]][2])
		return a > b
	})

	for i, key := range keys {

		newMapValue[key][3] = strconv.Itoa(i + 1)

	}
	return newMapValue
}
