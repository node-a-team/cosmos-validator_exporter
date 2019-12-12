package types

import ()

//validator status
type ValidatorStatus struct {
	Height int64
	Result []ValidatorStatusResult
}

type ValidatorStatusResult struct {
	Operator_address string `json:"operator_address"`
	Consensus_pubkey string `json:"consensus_pubkey"`
	Jailed           bool   `json:"jailed"`
	Status           int    `json:"status"`
	Tokens           string `json:"tokens"`
	Delegator_shares string `json:"delegator_shares"`
	Description      struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	}
	Unbonding_height string `json:"unbonding_height"`
	Unbonding_time   string `json:"unbonding_time"`
	Commission       struct {
		Commission_rates struct {
			Rate            string `json:"rate"`
			Max_rate        string `json:"max_rate"`
			Max_change_rate string `json:"max_change_rate"`
		}
		Update_time     string `json:"update_time"`
	}
	Min_self_delegation string `json:"min_self_delegation"`
}

// validator delegations
type ValidatorDelegationStatus struct {
	Height int64
        Result []ValidatorDelegationStatusResult
}
type ValidatorDelegationStatusResult struct {
	Delegator_Address string `json:"delegator_address"`
	Validator_Address string `json:"validator_address"`
	Shares            string `json:"shares"`
}


type ValidatorRewards struct {
	Height int64
        Result struct {
		Operator_address  string `json:"operator_address"`
		Self_Bond_Rewards []Coin `json:"self_bond_rewards"`
		Val_commission    []Coin `json:"val_commission"`
	}
}
