package cmd

type QueryContractMethodParams struct {
	FromAddress     string   `json:"from_address"`
	ContractAddress string   `json:"contract_address"`
	Method          string   `json:"method"`
	Args            []string `json:"args"`
	Abi             []byte   `json:"abi"`
}
