package types

import ()

type AccountStatus struct {
	Type  string `json:"type"`
	Value AccountValue
}

type BaseVestingAccountStatus struct {
	Type  string `json:"type"`
	Value struct {
		BaseVestingAccount struct {
			BaseAccount AccountValue
		}
	}
}

type AccountValue struct {
	Address			string `json:"address"`
	Coins			[]Coin `json:"coins"`
	Public_key struct {
		Type		string `json:"type"`
		Value		string `json:"value"`
	}
	Account_number		string `json:"account_number"`
	Sequence		string `json:"sequence"`
}
