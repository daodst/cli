package main

import (
	"freemasonry.cc/fm-cli/cmd"
	"freemasonry.cc/fm-cli/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func main() {
	setupConfig()
	rootCmd := &cobra.Command{
		Use:   "client",
		Short: "fm client util",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(cmd.CreateCmd())
	rootCmd.AddCommand(cmd.ViewCmd())
	rootCmd.AddCommand(cmd.TransferCmd())
	rootCmd.AddCommand(cmd.BalanceCmd())
	rootCmd.AddCommand(cmd.TxCmd())
	rootCmd.Execute()
}

func setupConfig() {
	config := sdk.GetConfig()
	core.SetBech32Prefixes(config)
	core.SetBip44CoinType(config)
	config.Seal()
}
