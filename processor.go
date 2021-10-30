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
		return errors.Wrap(err, "error opening file for reading")
	}
	defer inputHandle.close()

	outputHandle, err := NewFileOutput(t.outputFilePath)
	if err != nil {
		return errors.Wrap(err, "error opening file for writing")
	}
	defer outputHandle.Close()

	if err != nil {
		return errors.Wrap(err, "error opening input file")
	}
	// initilize mempool
	mp := mempool.NewMempool(t.capacity)
	//push all the transactions in input file to mempool
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

	// push all transactions in mempool to file
	for {
		txn := mp.Remove()
		//mempool is empty
		if (transaction.Transaction{}) == txn  {
			break;
		}
		err := outputHandle.WriteTransaction(txn)
		if err != nil {
			return errors.Wrap(err, "error reading transaction")
		}
	}

	return nil
}
