package minting

import (
	t "github.com/node-a-team/cosmos-validator_exporter/types"

	"encoding/json"
	"os/exec"
	"strconv"
)

var ()

func GetInflation() float64 {

	var inflation string
	var result float64

	cmd := "curl -s " + t.RestServer + "/minting/inflation"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &inflation)

	result, _ = strconv.ParseFloat(inflation, 64)

	return result

}

func MintingParamsStatus() t.MintingParams {

	var mintingParamsStatus t.MintingParams

	cmd := "curl -s " + t.RestServer + "/minting/parameters"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &mintingParamsStatus)

	return mintingParamsStatus

}
