package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/programcpp/kava-mempool/transaction"
)

type output interface {
	// Writes the next transaction to the output stream.
	WriteTransaction(t transaction.Transaction) error
	// Always cleanup once file IO is complete
	// returns error if already closed
	Close() error
}

// Note: do not use the struct whitout initialization.
// use NewFileOutput instead
type fileOutput struct {
	filePath string

	f *os.File
}

// write transaction to the file.
// appends the transaction, is the file already exists
func NewFileOutput(file string) (output, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fileOutput{}, errors.Wrap(err, "error opening file")
	}
	return fileOutput{
		filePath: file,
		f:        f,
	}, nil
}

func (f fileOutput) Close() error {
	err := f.f.Close()
	if err != nil {
		return errors.Wrap(err, "error closing file")
	}
	return nil
}

func (f fileOutput) WriteTransaction(t transaction.Transaction) error {
	formattedTransaction := fmt.Sprintf("TxHash=%s Gas=%d FeePerGas=%s Signature=%s\n",
		t.Hash, t.Gas, formatFeePerGas(t.FeePerGas), t.Signature)
	_, err := f.f.WriteString(formattedTransaction)
	if err != nil {
		return errors.Wrap(err, "error writing to file")
	}
	return nil
}

func formatFeePerGas(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
