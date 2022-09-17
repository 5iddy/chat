package keeper_test

import (
	"strconv"
	"testing"

	keepertest "chat/testutil/keeper"
	"chat/testutil/nullify"
	"chat/x/profile/keeper"
	"chat/x/profile/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNProfile(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Profile {
	items := make([]types.Profile, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetProfile(ctx, items[i])
	}
	return items
}

func TestProfileGet(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProfile(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProfileRemove(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProfile(ctx,
			item.Name,
		)
		_, found := keeper.GetProfile(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestProfileGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProfile(ctx)),
	)
}
