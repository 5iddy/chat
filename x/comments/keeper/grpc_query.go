package keeper

import (
	"chat/x/comments/types"
)

var _ types.QueryServer = Keeper{}
