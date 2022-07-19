package keeper

import (
	"zygosis/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
