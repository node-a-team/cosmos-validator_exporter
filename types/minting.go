package types

import (
)

type MintingParams struct {
	Mint_denom		string  `json:"mint_denom"`
	Inflation_rate_change	string	`json:"inflation_rate_change`
	Inflation_max		string	`json:"inflation_max"`
	Inflation_min		string	`json:"inflation_min"`
	Goal_bonded		string	`json:"goal_bonded"`
	Blocks_per_year		string	`json:"blocks_per_year"`
}

