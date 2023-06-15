package domain

type Transaction struct {
	Id         string
	DateTime   string
	Amount     string
	Instrument Instrument
	Contract   Contract
	Receivable Receivable
}
