package utils

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/tendermint/tendermint/libs/bech32"

	sdk "github.com/cosmos/cosmos-sdk/types"

)

const (
	bech32MainPrefix = "cosmos"
)

var (
	Bech32Prefixes = []string{
                // account's address
                bech32MainPrefix,
                // account's public key
                bech32MainPrefix+sdk.PrefixPublic,
                // validator's operator address
                bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator,
                // validator's operator public key
                bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic,
                // consensus node address
                bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus,
                // consensus node public key
                bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic,
        }
)

// Bech32 Addr -> Hex Addr
func Bech32AddrToHexAddr(bech32str string, log *zap.Logger) string {
	_, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "Bech32Addr To HexAddr"),)
        }

	return fmt.Sprintf("%X", bz)
}

func GetAccAddrFromOperAddr(operAddr string, log *zap.Logger) string {

        // Get HexAddress
        hexAddr, err := sdk.ValAddressFromBech32(operAddr)
	// log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "OperAddr To HexAddr"),)
        }

        accAddr, err := sdk.AccAddressFromHex(fmt.Sprint(hexAddr))
	// log
        if err != nil {
                // handle error
                log.Fatal("Utils-Address", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
//                log.Info("Utils-Address", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Change Address", "HexAddr To AccAddr"),)
        }

        return accAddr.String()
}

