package cmd

import (
	"encoding/json"
	"fmt"
	"freemasonry.cc/fm-cli/core"
	abi2 "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func ContractQueryCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "contract [contract_address] [from_address] [method] [abi]",
		Short: "contract query method",
		Args:  cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			contract := common.HexToAddress(args[0])
			from := common.HexToAddress(args[1])
			abi := abi2.ABI{}
			err := json.Unmarshal([]byte(args[3]), &abi)
			if err != nil {
				fmt.Println("error abi Unmarshal :", err.Error())
				return
			}
			param := cmd.Flag("args").Value.String()
			params := QueryContractMethodParams{FromAddress: from, ContractAddress: contract, Method: args[2], Args: param}
			route := fmt.Sprintf("custom/%s/%s", "contract", "contract_method")
			result, err := core.QueryWithData(route, params)
			if err != nil {
				fmt.Println("error QueryWithData :", err.Error())
				return
			}
			fmt.Println(string(result))
		},
	}
	return cmd
}
