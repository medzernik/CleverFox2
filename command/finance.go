package command

import (
	"github.com/almerlucke/go-iban/iban"
)

// Country identification
const (
	SVK = iota
	CZE
)

// This is the custom account number format (for now only czech)
type AccNumber struct {
	PreNumber int
	Number    int
	BankCode  int
}

func ParseIBAN(value string) (*iban.IBAN, error) {
	return iban.NewIBAN(value)
}

func IBANtoAccountNumber(value *iban.IBAN) (AccNumber, error) {

}
