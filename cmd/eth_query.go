package cmd

import (
	"encoding/hex"
	"fmt"
	"freemasonry.cc/cli/core"
	"github.com/spf13/cobra"
)

func EthQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [from_address] [to_address] [data]",
		Short: "contract query method",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			data, err := hex.DecodeString(args[2])
			if err != nil {
				fmt.Println("error DecodeString :", err.Error())
				return
			}
			rawUrl := cmd.Flag("rawUrl").Value.String()
			result, err := core.EthCallContract(args[0], args[1], rawUrl, data)
			if err != nil {
				fmt.Println("error QueryWithData :", err.Error())
				return
			}
			fmt.Println(result.String())
		},
	}
	cmd.Flags().StringP("rawUrl", "", "http://127.0.0.1:8545", "json rpc url")
	return cmd
}
