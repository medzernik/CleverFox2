package command

// Country identification
const (
	SVK = iota
	CZE
)

type IBAN struct {
	AccNumber string
}

type AccNumber struct {
	PreNumber int
	Number    int
	BankCode  int
}
