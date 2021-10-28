package main

func processTransactions(ioHandle io) {
	// 1. Read transactions from file
	// 2. Create priority mempool
	// 3. Write transactions to file
	ioHandle.readTransaction()

	ioHandle.writeTransaction(transaction{})
}
