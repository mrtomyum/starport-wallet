package keeper

import (
	"gitlab.com/mrtomyum/wallet/x/dex/types"
)

var _ types.QueryServer = Keeper{}
