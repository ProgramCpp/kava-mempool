package main

func main(){
	processTransactions(fileIO{
		inputFilePath: "transactions.txt",
		outputFilePath: "prioritized-transactions.txt",
	})
}

