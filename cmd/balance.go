package cmd

import (
	"encoding/json"
	"fmt"
	"freemasonry.cc/cli/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"
)

func BalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance [account]",
		Short: "query account balance",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			node := cmd.Flag("node").Value.String()
			if node != "" {
				core.SetRpc(node)
			}
			denom := core.ParseBaseCoin(cmd.Flag("denom").Value.String())
			balance, err := queryBalance(args[0], denom)
			if err != nil {
				fmt.Println("error query balance:", err.Error())
				return
			}
			balanceByte, err := json.Marshal(balance)
			if err != nil {
				fmt.Println("error Marshal:", err.Error())
				return
			}
			fmt.Println(string(balanceByte))
		},
	}
	cmd.Flags().StringP("denom", "", "", "Query specified currency")
	cmd.Flags().StringP("node", "", "", "node rpc url (default: tcp://127.0.0.1:26657)")
	return cmd
}

func queryBalance(address, denom string) (sdk.Coins, error) {
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdk.Coins{}, err
	}
	var (
		params interface{}
		route  string
	)
	if denom == "" {
		params = types.NewQueryAllBalancesRequest(addr, nil)
		route = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllBalances)
	} else {
		params = types.NewQueryBalanceRequest(addr, denom)
		route = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryBalance)
	}
	res, err := core.QueryWithData(route, params)
	if err != nil {
		return sdk.Coins{}, err
	}
	if denom == "" {
		coins, err := core.UnmarshlCoins(res)
		if err != nil {
			return sdk.Coins{}, err
		}
		return coins, nil
	} else {
		coin, err := core.UnmarshlCoin(res)
		if err != nil {
			return sdk.Coins{}, err
		}
		return sdk.NewCoins(coin), nil
	}
}
