package profile

import (
	"math/rand"

	"chat/testutil/sample"
	profilesimulation "chat/x/profile/simulation"
	"chat/x/profile/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = profilesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProfile = "op_weight_msg_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProfile int = 100

	opWeightMsgUpdateProfile = "op_weight_msg_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProfile int = 100

	opWeightMsgDeleteProfile = "op_weight_msg_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProfile int = 100

	opWeightMsgAddBioToProfile = "op_weight_msg_add_bio_to_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddBioToProfile int = 100

	opWeightMsgAddWebsiteToProfile = "op_weight_msg_add_website_to_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddWebsiteToProfile int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	profileGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ProfileList: []types.Profile{
			{
				Creator: sample.AccAddress(),
				Name:    "0",
			},
			{
				Creator: sample.AccAddress(),
				Name:    "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&profileGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProfile, &weightMsgCreateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProfile = defaultWeightMsgCreateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProfile,
		profilesimulation.SimulateMsgCreateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProfile, &weightMsgUpdateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProfile = defaultWeightMsgUpdateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProfile,
		profilesimulation.SimulateMsgUpdateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProfile, &weightMsgDeleteProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProfile = defaultWeightMsgDeleteProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProfile,
		profilesimulation.SimulateMsgDeleteProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddBioToProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddBioToProfile, &weightMsgAddBioToProfile, nil,
		func(_ *rand.Rand) {
			weightMsgAddBioToProfile = defaultWeightMsgAddBioToProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddBioToProfile,
		profilesimulation.SimulateMsgAddBioToProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddWebsiteToProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddWebsiteToProfile, &weightMsgAddWebsiteToProfile, nil,
		func(_ *rand.Rand) {
			weightMsgAddWebsiteToProfile = defaultWeightMsgAddWebsiteToProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddWebsiteToProfile,
		profilesimulation.SimulateMsgAddWebsiteToProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
