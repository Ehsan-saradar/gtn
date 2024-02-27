package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gtn/x/gtn/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	authority string,
	bankKeeper types.BankKeeper,
	logger log.Logger,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		bankKeeper:   bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ShareRewards(ctx sdk.Context, game types.Game, guesses []types.Guess) {
	if len(guesses) == 0 {
		// Transfer the reward back to the creator
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			sdk.MustAccAddressFromBech32(game.Creator),
			sdk.NewCoins(game.Reward),
		); err != nil {
			panic(err)
		}

		return
	}
	reward := game.Reward.Amount
	rewardShare := game.Reward.Amount.Quo(math.NewInt(int64(len(guesses))))

	for i, guess := range guesses {
		if i == len(guesses)-1 {
			rewardShare = reward
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			sdk.MustAccAddressFromBech32(guess.Participant),
			sdk.NewCoins(sdk.NewCoin(game.Reward.Denom, rewardShare), game.EntryFee),
		); err != nil {
			panic(err)
		}

		reward = reward.Sub(rewardShare)
	}
}

func (k Keeper) FundCreator(ctx sdk.Context, game types.Game, losersCount int) {
	// Transfer the entry fees to the creator
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		sdk.MustAccAddressFromBech32(game.Creator),
		sdk.NewCoins(sdk.NewCoin(game.EntryFee.Denom, game.EntryFee.Amount.Mul(math.NewInt(int64(losersCount))))),
	); err != nil {
		panic(err)
	}
}
