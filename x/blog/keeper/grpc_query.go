package keeper

import (
	"chat/x/blog/types"
)

var _ types.QueryServer = Keeper{}
