package command

import (
	"fmt"

	"github.com/almerlucke/go-iban/iban"
)

// This is the custom account number format (for now only czech)
type AccNumber struct {
	CountryCode string
	PreNumber   string
	Number      string
	BankCode    string
}

func (self *AccNumber) ParseToString() string {
	return fmt.Sprint(
		"Country code: ",
		self.CountryCode,
		"\n",
		"Account number: ",
		self.PreNumber,
		"-",
		self.Number,
		"/",
		self.BankCode,
	)
}

func ParseIBAN(value string) (*iban.IBAN, error) {
	return iban.NewIBAN(value)
}

func IBANtoAccountNumber(value *iban.IBAN) (AccNumber, error) {

	switch value.CountryCode {
	case "CZ":
		return (AccNumber{
			CountryCode: value.CountryCode,
			PreNumber:   value.BBAN[4:10],
			Number:      value.BBAN[10:],
			BankCode:    value.BBAN[0:4],
		}), nil

	default:
		return (AccNumber{}), fmt.Errorf("unsupported country")
	}

}
