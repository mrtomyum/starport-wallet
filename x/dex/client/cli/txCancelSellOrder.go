package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"gitlab.com/mrtomyum/wallet/x/dex/types"
)

var _ = strconv.Itoa(0)

func CmdCancelSellOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancelSellOrder [amountDenom] [orderID]",
		Short: "Broadcast message cancelSellOrder",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAmountDenom := string(args[0])
			argsOrderID, _ := strconv.ParseInt(args[1], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelSellOrder(clientCtx.GetFromAddress().String(), string(argsAmountDenom), int32(argsOrderID))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
