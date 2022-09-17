package keeper

import (
	"chat/x/profile/types"
)

var _ types.QueryServer = Keeper{}
