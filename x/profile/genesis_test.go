package profile_test

import (
	"testing"

	keepertest "chat/testutil/keeper"
	"chat/testutil/nullify"
	"chat/x/profile"
	"chat/x/profile/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProfileKeeper(t)
	profile.InitGenesis(ctx, *k, genesisState)
	got := profile.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
