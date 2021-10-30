package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessTransactions_ShouldSaveTransactionsAsPerPriority(t *testing.T) {
	testDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	testInputFile := filepath.Join(testDir, "testInput.txt")
	err = ioutil.WriteFile(testInputFile, []byte(
		"TxHash=1 Gas=3 FeePerGas=0.001 Signature=abc\n"+
			"TxHash=2 Gas=4 FeePerGas=0.001 Signature=def\n"+
			"TxHash=3 Gas=5 FeePerGas=0.001 Signature=ghi\n"+
			"TxHash=4 Gas=6 FeePerGas=0.001 Signature=jkl"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	outputTestFile := filepath.Join(testDir, "testOutput.txt")
	p := transactionProcessor{
		capacity:       3,
		inputFilePath:  testInputFile,
		outputFilePath: outputTestFile,
	}

	err = p.processTransactions()
	assert.NoError(t, err)

	output, err := ioutil.ReadFile(outputTestFile)
	assert.NoError(t, err)
	assert.Equal(t, []byte(
		"TxHash=2 Gas=4 FeePerGas=0.001 Signature=def\n"+
			"TxHash=3 Gas=5 FeePerGas=0.001 Signature=ghi\n"+
			"TxHash=4 Gas=6 FeePerGas=0.001 Signature=jkl\n"), output)
}
