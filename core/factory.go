package core

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	authType "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/ethermint/app"
	"github.com/evmos/ethermint/encoding"
	"github.com/spf13/pflag"
)

var clientCtx client.Context

var clientFactory tx.Factory

var encodingConfig params.EncodingConfig

func NewAccountClient() AccountClient {
	return AccountClient{NewSecretKey()}
}

func init() {
	encodingConfig = encoding.MakeConfig(app.ModuleBasics)

	clientCtx = client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithOffline(true).
		WithAccountRetriever(authType.AccountRetriever{})

	flags := pflag.NewFlagSet("chat", pflag.ContinueOnError)

	flagErrorBuf := new(bytes.Buffer)

	flags.SetOutput(flagErrorBuf)

	clientFactory = tx.NewFactoryCLI(clientCtx, flags)
	clientFactory.
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithTxConfig(clientCtx.TxConfig)
}
