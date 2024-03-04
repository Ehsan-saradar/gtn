package cli

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"gtn/x/gtn/types"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

var FlagOutput = "out"

// NewTxCmd returns a root CLI command handler for all x/gtn transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "GTN transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewCreateGame(),
	)

	return txCmd
}

// NewCreateGame returns a CLI command handler for creating a game and saving the salt and number in an encrypted file
// to be used in revealing the game later.
func NewCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-game [number] [duration] [entry-fee] [reward]",
		Short:   "Create a game and save the salt and number in an encrypted file",
		Example: fmt.Sprintf("%s tx gtn create-game 13 12500 10000stake 10stake", version.AppName),
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			number, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return errors.Wrap(err, "number")
			}

			duration, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return errors.Wrap(err, "duration")
			}

			entryFee, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return errors.Wrap(err, "entry fee")
			}

			reward, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return errors.Wrap(err, "reward")
			}

			salt := generateRandSalt()
			commitmentHash := types.CalculateCommitmentHash(salt, number)

			msg := types.NewMsgCreateGame(
				clientCtx.GetFromAddress().String(),
				commitmentHash,
				duration,
				entryFee,
				reward,
			)

			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if err != nil {
				return err
			}

			output, err := cmd.Flags().GetString(FlagOutput)
			if err != nil {
				return err
			}

			return saveSaltAndNumber(output, salt, number, commitmentHash)
		},
	}

	cmd.Flags().String(FlagOutput, "out.json", "The output file containing the encrypted salt and number")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func generateRandSalt() []byte {
	var salt [32]byte
	_, err := crand.Read(salt[:])
	if err != nil {
		panic(err)
	}

	return salt[:]
}

func saveSaltAndNumber(output string, salt []byte, number uint64, commitmentHash []byte) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	data := map[string]interface{}{
		"salt":            salt,
		"number":          number,
		"commitment_hash": commitmentHash,
	}
	return json.NewEncoder(f).Encode(data)
}
