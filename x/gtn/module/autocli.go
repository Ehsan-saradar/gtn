package gtn

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gtn/api/gtn/gtn"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "GameAll",
					Use:       "list-game",
					Short:     "List all game",
				},
				{
					RpcMethod:      "Game",
					Use:            "show-game [id]",
					Short:          "Shows a game by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "GuessAll",
					Use:       "list-guess",
					Short:     "List all guess",
				},
				{
					RpcMethod:      "GameGuesses",
					Use:            "list-game-guesses [game-id]",
					Short:          "List all guesses for a game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "game_id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateGame",
					Use:            "create-game [commitment-hash] [duration] [entry-fee] [reward]",
					Short:          "Create game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "commitment_hash"}, {ProtoField: "duration"}, {ProtoField: "entry_fee"}, {ProtoField: "reward"}},
				},
				{
					RpcMethod:      "RevealGame",
					Use:            "reveal-game [id] [salt] [number]",
					Short:          "Reveal game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "salt"}, {ProtoField: "number"}},
				},
				{
					RpcMethod:      "SubmitGuess",
					Use:            "submit-guess [game-id] [number]",
					Short:          "Submit guess",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "game_id"}, {ProtoField: "number"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
