package types

import (
	"time"
)

type CommitStatus struct {
	Result	CommitResult
}

type CommitResult struct {
	Signed_header	CommitSignedHeader
}

type CommitSignedHeader struct {
	Header	CommitHeader
	Commit	Commits
}

type CommitHeader struct {
	Chain_id		string
	Height			string
	Time			time.Time
	Num_txs			string
	Total_txs		string
	Proposer_address	string
}

type Commits struct {
	Precommits	[]Precommits
}

type Precommits struct {
	Type			int
	Height			string
	Round			string
	Timestamp		time.Time
	Validator_address	string
	Validator_index		string
	Signature		string
}
