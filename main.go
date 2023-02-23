package main

import (
	"freemasonry.cc/wallet/cmd"
	"freemasonry.cc/wallet/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func main() {
	setupConfig()
	rootCmd := &cobra.Command{
		Use:   "wallet",
		Short: "fm wallet util",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(cmd.CreateCmd())
	rootCmd.AddCommand(cmd.ViewCmd())
	rootCmd.Execute()
}

func setupConfig() {
	config := sdk.GetConfig()
	core.SetBech32Prefixes(config)
	core.SetBip44CoinType(config)
	config.Seal()
}
