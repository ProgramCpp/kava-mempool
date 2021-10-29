package main

func main(){
	transactionProcessor{
		capacity:       5000,
		inputFilePath:  "transactions.txt",
		outputFilePath: "prioritized-transactions.txt",
	}.processTransactions()
}

