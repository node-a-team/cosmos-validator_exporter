package keyutil

import (
        "encoding/hex"
        "fmt"

        t "github.com/node-a-team/cosmos-validator_exporter/types"
        "github.com/tendermint/tendermint/libs/bech32"
)

// Print info from HEX
func RunFromHex(hexaddr string) [6]string {

	var bech32Prefixes = []string{

	        // cosmos    ->    account's address
	        t.Bech32MainPrefix,
	        // cosmospub    ->    account's public key
	        t.Bech32MainPrefix + t.PrefixPublic,
	        // cosmosvaloper    ->    validator's operator address
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixOperator,
	        // cosmosvaloperpub    ->    validator's operator public key
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixOperator + t.PrefixPublic,
	        // cosmosvalcons    ->    consensus node address
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixConsensus,
	        // cosmosvalconspub    ->    consensus node public key
	        t.Bech32MainPrefix + t.PrefixValidator + t.PrefixConsensus + t.PrefixPublic,
	}


	// keys[0]: account's address
	// keys[1]: account's public key
	// keys[2]: validator's operator address
	// keys[3]: validator's operator public key
	// keys[4]: consensus node address
	// keys[5]: consensus node public key -> No tendermint show-validator
	var keys [6]string

        bz, _ := hex.DecodeString(hexaddr)

        for i, prefix := range bech32Prefixes {
                bech32Addr, err := bech32.ConvertAndEncode(prefix, bz)

                if err != nil {
                        panic(err)
                }

		keys[i] = bech32Addr
        }

	return keys
}


// Print info from bech32.
func RunFromBech32(bech32str string) string{
        _, bz, err := bech32.DecodeAndConvert(bech32str)
        if err != nil {
                fmt.Println("Not a valid bech32 string")
                return "function/keyutil) RunFrombech32() Err"
        }

        return fmt.Sprintf("%X", bz)
}


// Operator address -> Ohter address
func OperAddrToOtherAddr(operaddr string) [6]string {

        hexOperaddr := RunFromBech32(operaddr)
        keys := RunFromHex(hexOperaddr)

        return keys
}

