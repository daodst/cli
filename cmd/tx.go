package cmd

import (
	"encoding/json"
	"fmt"
	"freemasonry.cc/fm-cli/core"
	"github.com/spf13/cobra"
)

func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx [hash]",
		Short: "query tx status",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			node := cmd.Flag("node").Value.String()
			if node != "" {
				core.SetRpc(node)
			}
			result, notFound, err := core.QueryTx(args[0])
			if err != nil {
				fmt.Println("-----------------------------------------------------------")
				fmt.Println("error query hash:", err.Error())
				return
			}
			if notFound {
				fmt.Println("-----------------------------------------------------------")
				fmt.Println("error tx hash not found ! ! !")
				return
			}
			resultByte, _ := json.Marshal(result)
			fmt.Println(string(resultByte))
		},
	}
	cmd.Flags().StringP("node", "", "", "node rpc url (default: tcp://127.0.0.1:26657)")
	return cmd
}
