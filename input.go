package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/programcpp/kava-mempool/transaction"
)

type input interface {
	// Reads the next transaction from the input stream.
	// to start reading from the beginning, create a new instance
	// returns empty transaction when EOF is reached
	readTransaction() (transaction.Transaction, error)
	// Always cleanup once file IO is complete
	// returns error if already closed
	close() error
}

// Note: do not use the struct whitout initialization.
// use NewFileInput instead
type fileInput struct {
	inputFilePath string

	f *os.File
	sc *bufio.Scanner
}

func NewFileInput(iFile string) (input, error) {
	f, err := os.Open(iFile)
	if err != nil {
		return nil, errors.Wrap(err, "error opening file")
	}
	return &fileInput{
		inputFilePath: iFile,
		f:             f,
		sc:            bufio.NewScanner(f),
	}, nil
}

func (f fileInput) close() error {
	err := f.f.Close()
	if err != nil {
		return errors.Wrap(err, "error closing file")
	}
	return nil
}


func (f fileInput) readTransaction() (transaction.Transaction, error) {
	f.sc.Scan()
	if f.sc.Err() != nil {
		return transaction.Transaction{}, errors.Wrap(f.sc.Err(), "error reading file")
	}
	var txn transaction.Transaction
	fmt.Sscanf(f.sc.Text(), "TxHash=%s Gas=%d FeePerGas=%f Signature=%s", &txn.Hash, &txn.Gas, &txn.FeePerGas, &txn.Signature)
	return txn, nil
}