package command

import (
	"github.com/almerlucke/go-iban/iban"
)

// Country identification
const (
	SVK = iota
	CZE
)

type IntBankNUM struct {
	AccNumber string
}

type AccNumber struct {
	PreNumber int
	Number    int
	BankCode  int
}

func ParseIBAN(value string) (*iban.IBAN, error) {
	return iban.NewIBAN(value)
}
