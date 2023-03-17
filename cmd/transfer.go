package cmd

import (
	"encoding/json"
	"fmt"
	"freemasonry.cc/cli/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"
	"strings"
)

func TransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [pk] [from] [to] [amount]",
		Short: "transfer",
		Long: fmt.Sprintf(`Example:
		cli transfer 0B9A99EDA4B4C12323C4C1EB1210E98E8FE22FE1AE3543CACD044825CC9A793B  dex10436g09cfa2nz2k6yphzj9qs9u93qtvvkkftur  dex1va8aaeystat4twpy70ns7235pxwczg0698twv4  100dst,100fm
		
		transfer amount (eg:1dst or 1dst,1fm)
`),
		Args: cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			node := cmd.Flag("node").Value.String()
			if node != "" {
				core.SetRpc(node)
			}
			memo := cmd.Flag("memo").Value.String()
			resp, err := transfer(args[1], args[2], args[3], memo, args[0])
			if err != nil {
				fmt.Println("error transfer:", err.Error())
				return
			}

			respData, err := json.Marshal(resp)
			if err != nil {
				fmt.Println("error Marshal:", err.Error())
				return
			}
			fmt.Println(string(respData))
		},
	}
	cmd.Flags().StringP("memo", "", "", "transfer description")
	cmd.Flags().StringP("node", "", "", "node rpc url (default: tcp://127.0.0.1:26657)")
	return cmd
}

func transfer(from, to, amount, memo, privateKey string) (*sdk.TxResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return nil, err
	}
	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return nil, err
	}
	seq, err := core.GetAccountNumberSequence(from)
	if err != nil {
		return nil, err
	}
	msg := bankTypes.NewMsgSend(fromAddr, toAddr, SplitAmount(amount))
	signedTx, err := core.SignTx(privateKey, seq, core.DefaultFee, memo, msg)
	if err != nil {
		return nil, err
	}
	signedTxBytes, err := core.SignTx2Bytes(signedTx)
	if err != nil {
		return nil, err
	}
	res, err := core.BroadcastTx(signedTxBytes)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SplitAmount(amount string) sdk.Coins {
	coins := sdk.Coins{}
	amStr := strings.Split(amount, ",")
	for _, am := range amStr {
		denom := core.DisplayDenom
		if strings.Contains(am, core.GovDenom) {
			denom = core.GovDenom
		}
		a := strings.Split(strings.ToLower(am), denom)
		coin := core.RealString2LedgerCoin(a[0], denom)
		coins = append(coins, coin)
	}
	return coins
}
