package command

import (
	"CleverFox2/logging"
	"CleverFox2/tviewsystem"
	"bytes"
	"compress/lzw"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
	"time"

	"image/png"
	"os"

	"github.com/almerlucke/go-iban/iban"

	"github.com/boombuler/barcode/qr"
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

type Payment struct {
	Amount           float64
	IBAN             string
	Swift            string
	Date             time.Time
	BeneficiaryName  string
	Currency         string
	VariableSymbol   string
	ConstantSymbol   string
	SpecificSymbol   string
	Note             string
	BeneficiaryAddr1 string
	BeneficiaryAddr2 string
}

func (self *Payment) ParseToString() string {
	// if self.Date == nil {
	// 	self.Date = time.Now()
	// }
	return strings.Join([]string{
		"",
		"1", // payment
		"1", // simple payment
		fmt.Sprintf("%.2f", self.Amount),
		self.Currency,
		self.Date.Format("2006-01-02 15:04:05"),
		self.VariableSymbol,
		self.ConstantSymbol,
		self.SpecificSymbol,
		"", // previous 3 entries in SEPA format, empty because already provided above
		self.Note,
		"1", // to an account
		self.IBAN,
		self.Swift,
		"0", // not recurring
		"0", // not 'inkaso'
		self.BeneficiaryName,
		self.BeneficiaryAddr1,
		self.BeneficiaryAddr2,
	}, "\t")
}

func GenerateQRCode() {

	// Generating the data temporarily
	payment := Payment{
		Amount:           1.0,
		IBAN:             "SK7700000000000000000000",
		Swift:            "FIOZSKBAXXX",
		Date:             time.Now(),
		BeneficiaryName:  "Foo",
		Currency:         "EUR",
		VariableSymbol:   "11",
		ConstantSymbol:   "22",
		SpecificSymbol:   "33",
		Note:             "bar",
		BeneficiaryAddr1: "address 1",
		BeneficiaryAddr2: "address 2",
	}

	logging.Log.Info(payment.ParseToString())

	// Checksum calculation
	checksum := crc32.ChecksumIEEE([]byte(payment.ParseToString()))
	checksumBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(checksumBytes, checksum)

	total := append(checksumBytes, []byte(payment.ParseToString())...)

	// This is the XZ compression
	var buf bytes.Buffer
	w := lzw.NewWriter(&buf, lzw.LSB, 8)
	w.Write(total)
	w.Close()
	compressed := buf.Bytes()

	// Prepends the length and converts to hexadecimal
	compressedWithLength := make([]byte, 4+len(compressed))
	binary.LittleEndian.PutUint16(compressedWithLength[2:], uint16(len(total)))
	copy(compressedWithLength[4:], compressed)
	compressedWithLengthHex := hex.EncodeToString(compressedWithLength)

	// Converst to a padded hex string
	binaryString := ""
	for _, b := range compressedWithLengthHex {
		binaryString += fmt.Sprintf("%08b", b)
	}

	// Pad zeroes on the right with multiples of 5
	length := len(binaryString)
	remainder := length % 5
	if remainder != 0 {
		binaryString += strings.Repeat("0", 5-remainder)
		length += 5 - remainder
	}

	subst := "0123456789ABCDEFGHIJKLMNOPQRSTUV"
	// subst := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"
	result := ""
	for i := 0; i < len(binaryString); i += 5 {
		quintet, _ := strconv.ParseInt(binaryString[i:i+5], 2, 0)
		result += string(subst[quintet])
	}

	tviewsystem.StatusPush(result)

	qrCode, _ := qr.Encode(result, qr.L, qr.Auto)

	// Scale the barcode to 200x200 pixels
	// qrCode, _ = barcode.Scale(qrCode, 600, 600)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
