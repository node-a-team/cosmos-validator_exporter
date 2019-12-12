package types

import (
)

var (
)

type ValidatorsetsStatus struct {
	Height		int64
	Result struct {
		Block_height	string
		Validators	[]Validators
	}
}

type Validators struct {

	Address	string
	Pub_key	string
	Proposer_priority	string
	Voting_power	string
}

