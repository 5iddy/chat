package keeper_test

import (
	"context"
	"testing"

	keepertest "chat/testutil/keeper"
	"chat/x/comments/keeper"
	"chat/x/comments/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CommentsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
