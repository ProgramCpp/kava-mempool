package main

type io interface {
	writeTransactions([]transaction)
	readTransactions()[]transaction
}
type fileIO struct {
	inputFilePath string
	outputFilePath string
}

func (f fileIO) writeTransactions(_ []transaction) {
	panic("unimplemented")
}

func (f fileIO) readTransactions() []transaction {
	panic("unimplemented")
}
