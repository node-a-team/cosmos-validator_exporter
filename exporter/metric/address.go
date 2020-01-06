package metric

import (

	"fmt"
//	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"

//	"github.com/node-a-team/terra-validator_exporter/getData"
//	cfg "github.com/node-a-team/terra-validator_exporter/config"
//	utils "github.com/node-a-team/terra-validator_exporter/utils"
)

var (

)

func GetAccAddrFromOperAddr(operAddr string) string {

	// Get HexAddress
	hexAddr, err := sdk.ValAddressFromBech32(operAddr)
	if err != nil {
		// Error
	}

	accAddr, err := sdk.AccAddressFromHex(fmt.Sprint(hexAddr))
        if err != nil {
                // Error
        }

	return accAddr.String()
}
