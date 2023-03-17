package cmd

import (
	"encoding/json"
	"fmt"
	"freemasonry.cc/cli/core"
	"github.com/spf13/cobra"
)

func EthTransactionByHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash [hash]",
		Short: "query Transaction by hash",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rawUrl := cmd.Flag("rawUrl").Value.String()
			result, err := core.EthTransactionReceipt(args[0], rawUrl)
			if err != nil {
				fmt.Println("error QueryWithData :", err.Error())
				return
			}
			resultByte, err := json.Marshal(result)
			if err != nil {
				fmt.Println("error Marshal result:", err.Error())
				return
			}
			fmt.Println(string(resultByte))
		},
	}
	cmd.Flags().StringP("rawUrl", "", "http://127.0.0.1:8545", "json rpc url")
	return cmd
}
