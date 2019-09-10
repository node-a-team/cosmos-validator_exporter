package types

import (
	"time"
)

type ConsensusStatus struct {
	Jsonrpc					string		`json:"jsonrpc"`
	Result struct {
		Round_state struct {
			Status			string		`json:"height/round/step"`
			Start_time		time.Time
			Proposer_block_hash	string
			Locked_block_hash	string
			Valid_block_hash	string
			Height_vote_set		[]Height_vote_set
		}
	}
}

type Height_vote_set struct {
	Round			string
	Prevotes		[]string
	Prevotes_bit_array	string
	Precommits		[]string
	Precommits_bit_array	string
}

