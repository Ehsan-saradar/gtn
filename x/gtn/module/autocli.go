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
					Use:            "create-game [commitmentHash] [duration] [entryFee] [reward]",
					Short:          "Create game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "commitmentHash"}, {ProtoField: "duration"}, {ProtoField: "entryFee"}, {ProtoField: "reward"}},
				},
				{
					RpcMethod:      "UpdateGame",
					Use:            "update-game [id] [commitmentHash] [duration] [entryFee] [reward]",
					Short:          "Update game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "commitmentHash"}, {ProtoField: "duration"}, {ProtoField: "entryFee"}, {ProtoField: "reward"}},
				},
				{
					RpcMethod:      "DeleteGame",
					Use:            "delete-game [id]",
					Short:          "Delete game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
