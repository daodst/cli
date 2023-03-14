package cmd

import (
	"fmt"
	"freemasonry.cc/fm-cli/core"
	"github.com/spf13/cobra"
)

func ViewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "view account",
		Run: func(cmd *cobra.Command, args []string) {
			accC := core.NewAccountClient()

			seed := cmd.Flag("seed").Value.String() //dpos
			privateKey := cmd.Flag("pk").Value.String()

			if seed == "" && privateKey == "" {
				cmd.Help()
				fmt.Println("-----------------------------------------------------------")
				fmt.Println("error:", "Please enter a mnemonic or private key")
				return
			}

			if seed != "" && privateKey != "" {
				cmd.Help()
				fmt.Println("-----------------------------------------------------------")
				fmt.Println("error:", "Only one mnemonic and private key can be selected")
				return
			}

			if seed != "" {
				fmt.Printf("Mnemonic words:\t %s \n", seed)
				wallet, err := accC.CreateAccountFromSeed(seed)
				if err != nil {
					fmt.Println("error:", err.Error())
					return
				}
				wallet.Print()
			}

			if privateKey != "" {
				wallet, err := accC.CreateAccountFromPriv(privateKey)
				if err != nil {
					fmt.Println("error:", err.Error())
					return
				}
				wallet.Print()
			}
		},
	}
	cmd.Flags().StringP("seed", "", "", "mnemonic words")
	cmd.Flags().StringP("pk", "", "", "private key")
	return cmd
}
