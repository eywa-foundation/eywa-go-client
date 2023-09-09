package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-user [pubkey]",
		Short: "Broadcast message register-user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPubkey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterUser(
				clientCtx.GetFromAddress().String(),
				argPubkey,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
