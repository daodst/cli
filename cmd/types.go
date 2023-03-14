package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type QueryContractMethodParams struct {
	FromAddress     common.Address `json:"from_address"`
	ContractAddress common.Address `json:"contract_address"`
	Method          string         `json:"method"`
	Args            interface{}    `json:"args"`
	Abi             abi.ABI        `json:"abi"`
}
