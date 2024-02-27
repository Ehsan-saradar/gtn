package gtn

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"gtn/testutil/sample"
	gtnsimulation "gtn/x/gtn/simulation"
	"gtn/x/gtn/types"
)

// avoid unused import issue
var (
	_ = gtnsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGame int = 100

	opWeightMsgUpdateGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGame int = 100

	opWeightMsgDeleteGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGame int = 100

	opWeightMsgCreateGuess = "op_weight_msg_guess"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGuess int = 100

	opWeightMsgUpdateGuess = "op_weight_msg_guess"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGuess int = 100

	opWeightMsgDeleteGuess = "op_weight_msg_guess"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGuess int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gtnGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		GameList: []types.Game{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		GameCount: 2,
		GuessList: []types.Guess{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		GuessCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gtnGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGame int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGame, &weightMsgCreateGame, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGame = defaultWeightMsgCreateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGame,
		gtnsimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateGame int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateGame, &weightMsgUpdateGame, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGame = defaultWeightMsgUpdateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateGame,
		gtnsimulation.SimulateMsgUpdateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGame int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteGame, &weightMsgDeleteGame, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGame = defaultWeightMsgDeleteGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGame,
		gtnsimulation.SimulateMsgDeleteGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateGuess int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGuess, &weightMsgCreateGuess, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGuess = defaultWeightMsgCreateGuess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGuess,
		gtnsimulation.SimulateMsgCreateGuess(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateGuess int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateGuess, &weightMsgUpdateGuess, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGuess = defaultWeightMsgUpdateGuess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateGuess,
		gtnsimulation.SimulateMsgUpdateGuess(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGuess int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteGuess, &weightMsgDeleteGuess, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGuess = defaultWeightMsgDeleteGuess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGuess,
		gtnsimulation.SimulateMsgDeleteGuess(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGame,
			defaultWeightMsgCreateGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateGame,
			defaultWeightMsgUpdateGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgUpdateGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteGame,
			defaultWeightMsgDeleteGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgDeleteGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGuess,
			defaultWeightMsgCreateGuess,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgCreateGuess(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateGuess,
			defaultWeightMsgUpdateGuess,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgUpdateGuess(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteGuess,
			defaultWeightMsgDeleteGuess,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgDeleteGuess(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
