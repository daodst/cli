package core

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txSigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	evmhd "github.com/evmos/ethermint/crypto/hd"
	"github.com/spf13/pflag"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"regexp"
	"strings"
)

type AccountNumberSeqResponse struct {
	AccountNumber uint64 `json:"account_number"`
	Sequence      uint64 `json:"sequence"`
	NotFound      bool   `json:"not_found"`
}

func QueryWithData(route string, params interface{}) ([]byte, error) {
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	res, _, err := clientCtx.QueryWithData(route, bz)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryTx(tx_hash string) (resultTx *ctypes.ResultTx, notFound bool, err error) {
	txhash, err := hex.DecodeString(tx_hash)
	if err != nil {
		return
	}
	node, err := clientCtx.GetNode()
	if err != nil {
		return
	}
	resultTx, err = node.Tx(context.Background(), txhash, true)
	if err != nil {
		notFound = isTxNotFoundError(err.Error())
		if notFound {
			err = nil
		}
		return
	}
	return
}

func UnmarshlCoins(data []byte) (sdk.Coins, error) {
	coins := sdk.Coins{}
	err := clientCtx.LegacyAmino.UnmarshalJSON(data, &coins)
	if err != nil {
		return coins, err
	}
	return coins, nil
}
func UnmarshlCoin(data []byte) (sdk.Coin, error) {
	coin := sdk.Coin{}
	err := clientCtx.LegacyAmino.UnmarshalJSON(data, &coin)
	if err != nil {
		return coin, err
	}
	return coin, nil
}

func BroadcastTx(txBytes []byte) (*sdk.TxResponse, error) {
	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetAccountNumberSequence(address string) (AccountNumberSeqResponse, error) {
	seqRes := AccountNumberSeqResponse{AccountNumber: 0, Sequence: 0}
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return seqRes, err
	}
	accountNumber, sequence, err := clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, addr)
	if err != nil {
		if strings.Contains(err.Error(), "not found: key not found") {
			seqRes.NotFound = true
			return seqRes, nil
		}
		return seqRes, err
	}
	seqRes.Sequence = sequence
	seqRes.AccountNumber = accountNumber
	return seqRes, nil
}

func SignTx(privateKey string, seqDetail AccountNumberSeqResponse, fee legacytx.StdFee, memo string, msgs ...sdk.Msg) (signing.Tx, error) {
	flags := pflag.NewFlagSet("chat", pflag.ContinueOnError)
	flagErrorBuf := new(bytes.Buffer)
	flags.SetOutput(flagErrorBuf)
	clientFactory = tx.NewFactoryCLI(clientCtx, flags)
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	keyringAlgos := keyring.SigningAlgoList{evmhd.EthSecp256k1}
	algo, err := keyring.NewSigningAlgoFromString("eth_secp256k1", keyringAlgos)
	if err != nil {
		return nil, err
	}
	privKey := algo.Generate()(privKeyBytes)

	signMode := clientCtx.TxConfig.SignModeHandler().DefaultMode()
	signerData := signing.SignerData{
		ChainID:       clientCtx.ChainID,
		AccountNumber: seqDetail.AccountNumber,
		Sequence:      seqDetail.Sequence,
	}
	txBuild, err := tx.BuildUnsignedTx(clientFactory, msgs...)
	if err != nil {
		return nil, err
	}
	txBuild.SetGasLimit(fee.Gas)
	txBuild.SetFeeAmount(fee.Amount)
	txBuild.SetMemo(memo)
	sigData := txSigning.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := txSigning.SignatureV2{
		PubKey:   privKey.PubKey(),
		Data:     &sigData,
		Sequence: seqDetail.Sequence,
	}
	if err := txBuild.SetSignatures(sig); err != nil {
		return nil, err
	}
	signV2, err := tx.SignWithPrivKey(signMode, signerData, txBuild, privKey, clientCtx.TxConfig, seqDetail.Sequence)
	if err != nil {
		return nil, err
	}
	err = txBuild.SetSignatures(signV2)
	if err != nil {
		return nil, err
	}
	signedTx := txBuild.GetTx()
	return signedTx, nil
}

func SignTx2Bytes(signTxs signing.Tx) ([]byte, error) {
	return clientCtx.TxConfig.TxEncoder()(signTxs)
}

func isTxNotFoundError(errContent string) (ok bool) {
	errRegexp := `tx\ \([0-9A-Za-z]{64}\)\ not\ found`
	r, err := regexp.Compile(errRegexp)
	if err != nil {
		return false
	}
	if r.Match([]byte(errContent)) {
		return true
	} else {
		return false
	}
}
