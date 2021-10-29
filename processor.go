package main

import (
	"github.com/pkg/errors"
	"github.com/programcpp/kava-mempool/mempool"
	"github.com/programcpp/kava-mempool/transaction"
)

type transactionProcessor struct {
	capacity int
	inputFilePath string
	outputFilePath string
}

func (t transactionProcessor) processTransactions() error {
	// 1. initialize IO handles
	// 2. initialize mempool
	// 3. Read transactions from file
	// 4. Create priority mempool
	// 5. Write sorted transactions to file
	inputHandle, err := NewFileInput(t.inputFilePath)
	if err != nil {
		return errors.Wrap(err, "error opening input file")
	}
	mp := mempool.NewMempool(t.capacity)
	for {
		txn, err := inputHandle.readTransaction()
		if err != nil {
			return errors.Wrap(err, "error reading transaction")
		}
		// reached EOF
		if (transaction.Transaction{}) == txn  {
			break;
		}
		mp.Insert(txn)
	}

	return nil
}
