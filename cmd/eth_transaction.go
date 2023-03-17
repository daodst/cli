package cmd

import (
	"fmt"
	"freemasonry.cc/cli/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"math/big"
	"strconv"
)

func EthTransactionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction [from_address] [to_address] [data] [pk]",
		Short: "eth transaction",
		Long:  fmt.Sprintf(`notices:data Do not start with 0x`),
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			node := cmd.Flag("node").Value.String()
			if node != "" {
				core.SetRpc(node)
			}
			rawUrl := cmd.Flag("rawUrl").Value.String()
			gasLimitStr := cmd.Flag("gasLimit").Value.String()
			gasLimit, err := strconv.ParseUint(gasLimitStr, 10, 64)
			if err != nil {
				fmt.Println("error ParseUint:", err.Error())
				return
			}
			chainId, err := core.EthChainId(rawUrl)
			if err != nil {
				fmt.Println("error EthChainId:", err.Error())
				return
			}
			ethFromAddr := common.HexToAddress(args[0])
			cosmosAddr := sdk.AccAddress(ethFromAddr.Bytes())
			seq, err := core.GetAccountNumberSequence(cosmosAddr.String())
			if err != nil {
				fmt.Println("error GetAccountNumberSequence:", err.Error())
				return
			}
			gasPrice, err := core.EthGasPrice(rawUrl)
			if err != nil {
				fmt.Println("error EteGasPrice:", err.Error())
				return
			}
			core.LocalSendTx(seq.Sequence, gasLimit, big.NewInt(0), gasPrice, chainId, args[1], args[2], args[3], rawUrl)
		},
	}
	cmd.Flags().StringP("node", "", "", "node rpc url (default: tcp://127.0.0.1:26657)")
	cmd.Flags().StringP("rawUrl", "", "http://127.0.0.1:8545", "json rpc url")
	cmd.Flags().Uint64("gasLimit", 210000, "gasLimit")
	return cmd
}
