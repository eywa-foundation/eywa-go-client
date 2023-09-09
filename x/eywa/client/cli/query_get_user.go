package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-user [submitter]",
		Short: "Query get-user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqSubmitter := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetUserRequest{

				Submitter: reqSubmitter,
			}

			res, err := queryClient.GetUser(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
