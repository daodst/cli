package core

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	authType "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/ethermint/app"
	"github.com/evmos/ethermint/encoding"
	"github.com/spf13/pflag"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

var clientCtx client.Context

var clientFactory tx.Factory

var encodingConfig params.EncodingConfig

func NewAccountClient() AccountClient {
	return AccountClient{NewSecretKey()}
}

func SetRpc(rpc string) {
	rpcClient, err := rpchttp.New(rpc, "/websocket")
	if err != nil {
		panic("start ctx client error.")
	}
	clientCtx = clientCtx.WithNodeURI(rpc).WithClient(rpcClient)
}

func init() {
	encodingConfig = encoding.MakeConfig(app.ModuleBasics)
	rpcClient, err := rpchttp.New(RpcURL, "/websocket")
	if err != nil {
		panic("start ctx client error.")
	}
	clientCtx = client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithOffline(true).
		WithNodeURI(RpcURL).
		WithChainID(ChainID).
		WithClient(rpcClient).
		WithBroadcastMode(flags.BroadcastSync).
		WithAccountRetriever(authType.AccountRetriever{})

	flags := pflag.NewFlagSet("chat", pflag.ContinueOnError)

	flagErrorBuf := new(bytes.Buffer)

	flags.SetOutput(flagErrorBuf)

	clientFactory = tx.NewFactoryCLI(clientCtx, flags)
	clientFactory.
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithTxConfig(clientCtx.TxConfig)
}
