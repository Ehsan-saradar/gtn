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

	opWeightMsgRevealGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealGame int = 100

	opWeightMsgSubmitGuess = "op_weight_msg_guess"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitGuess int = 100

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

	var weightMsgRevealGame int
	simState.AppParams.GetOrGenerate(opWeightMsgRevealGame, &weightMsgRevealGame, nil,
		func(_ *rand.Rand) {
			weightMsgRevealGame = defaultWeightMsgRevealGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealGame,
		gtnsimulation.SimulateMsgRevealGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitGuess int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitGuess, &weightMsgSubmitGuess, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitGuess = defaultWeightMsgSubmitGuess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitGuess,
		gtnsimulation.SimulateMsgCreateGuess(am.accountKeeper, am.bankKeeper, am.keeper),
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
			opWeightMsgRevealGame,
			defaultWeightMsgRevealGame,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgRevealGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitGuess,
			defaultWeightMsgSubmitGuess,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gtnsimulation.SimulateMsgCreateGuess(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
