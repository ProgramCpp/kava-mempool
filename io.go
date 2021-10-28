package main

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

type io interface {
	writeTransaction(transaction)
	readTransaction() (transaction, error)
}
type fileIO struct {
	inputFilePath  string
	outputFilePath string
}

func (f fileIO) writeTransaction(_ transaction) {
	panic("unimplemented")
}

func (f fileIO) readTransaction() (transaction, error) {
	file, err := os.Open(f.inputFilePath)
	if err != nil {
		return transaction{}, errors.Wrap(err, "error opening file")
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	is_eof := sc.Scan()
	if sc.Err() != nil {
		return transaction{}, errors.Wrap(sc.Err(), "error reading file")
	}
	if is_eof{
		return transaction{}, nil
	}
	return transaction{}, nil	
}
