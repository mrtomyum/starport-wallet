package keeper

import (
	"gitlab.com/mrtomyum/wallet/x/wallet/types"
)

var _ types.QueryServer = Keeper{}
