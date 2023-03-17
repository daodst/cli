package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func EthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eth",
		Short: "eth subcommands",
		RunE:  client.ValidateCmd,
	}
	cmd.AddCommand(
		EthQueryCmd(),
		EthTransactionCmd(),
		EthTransactionByHash(),
	)
	return cmd
}
