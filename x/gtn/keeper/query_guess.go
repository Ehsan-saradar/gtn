package keeper

import (
	"context"

	"gtn/x/gtn/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GuessAll(ctx context.Context, req *types.QueryAllGuessRequest) (*types.QueryAllGuessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var guesss []types.Guess

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	guessStore := prefix.NewStore(store, types.KeyPrefix(types.GuessKey))

	pageRes, err := query.Paginate(guessStore, req.Pagination, func(key []byte, value []byte) error {
		var guess types.Guess
		if err := k.cdc.Unmarshal(value, &guess); err != nil {
			return err
		}

		guesss = append(guesss, guess)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGuessResponse{Guesses: guesss, Pagination: pageRes}, nil
}

func (k Keeper) GameGuesses(ctx context.Context, req *types.QueryGetGameGuessesRequest) (*types.QueryGetGameGuessesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	guesses := k.GetGameGuesses(ctx, req.GameId)

	return &types.QueryGetGameGuessesResponse{Guesses: guesses}, nil
}
