package domain

type Contract struct {
	ContractVersion    string
	CustomerId         string
	Mdr                string
	AllowedInstruments []string
	SettlementAddress  SettlementAddress
}
