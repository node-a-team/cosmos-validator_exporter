package metric

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
