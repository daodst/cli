package core

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

type RealCoins []RealCoin

type RealCoin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

var (
	MinRealAmountDec = sdk.NewDecWithPrec(1, 10)

	LedgerToRealRate = sdk.NewDecWithPrec(1, 18)

	RealToLedgerRateDec = sdk.MustNewDecFromStr("1000000000000000000")

	DefaultFee = legacytx.NewStdFee(flags.DefaultGasLimit, sdk.NewCoins(sdk.NewCoin(BaseDenom, sdk.MustNewDecFromStr("100000000000000000").TruncateInt())))
)

func ParseBaseCoin(symbol string) string {
	if symbol == DisplayDenom {
		return BaseDenom
	}
	return symbol
}

func RealString2LedgerCoin(realString, denom string) sdk.Coin {
	denom = ParseBaseCoin(denom)
	return sdk.NewCoin(denom, MustRealString2LedgerInt(realString))
}

func MustLedgerCoins2RealCoins(ledgerCoins sdk.Coins) (realCoins RealCoins) {
	for i := 0; i < len(ledgerCoins); i++ {
		realCoins = append(realCoins, MustLedgerCoin2RealCoin(ledgerCoins[i]))
	}
	return
}

func MustLedgerCoin2RealCoin(ledgerCoin sdk.Coin) (realCoin RealCoin) {
	ledgerAmount := ledgerCoin.Amount.ToDec()
	return RealCoin{
		Denom:  ledgerCoin.Denom,
		Amount: ledgerAmount.Mul(LedgerToRealRate).String(),
	}
}

func MustRealString2LedgerInt(realString string) (ledgerInt sdk.Int) {
	realIntAmountDec := sdk.MustNewDecFromStr(realString)
	if realIntAmountDec.LT(MinRealAmountDec) && !realIntAmountDec.IsZero() {
		realIntAmountDec = MinRealAmountDec
	}
	return realIntAmountDec.Mul(RealToLedgerRateDec).TruncateInt()
}
