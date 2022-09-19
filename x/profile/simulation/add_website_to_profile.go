package simulation

import (
	"math/rand"

	"chat/x/profile/keeper"
	"chat/x/profile/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddWebsiteToProfile(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddWebsiteToProfile{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddWebsiteToProfile simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddWebsiteToProfile simulation not implemented"), nil, nil
	}
}
