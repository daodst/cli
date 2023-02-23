package cmd

import (
	"fmt"
	"freemasonry.cc/wallet/core"
	"github.com/spf13/cobra"
)

func CreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create new account",
		Run: func(cmd *cobra.Command, args []string) {
			accC := core.NewAccountClient()
			seed, err := accC.CreateSeedWord()
			if err != nil {
				fmt.Println("error:", err.Error())
				return
			}
			fmt.Printf("Mnemonic words:\t %s \n", seed)
			wallet, err := accC.CreateAccountFromSeed(seed)
			if err != nil {
				fmt.Println("error:", err.Error())
				return
			}
			wallet.Print()
		},
	}
	return cmd
}
