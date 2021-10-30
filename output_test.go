package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/programcpp/kava-mempool/transaction"
	"github.com/stretchr/testify/assert"
)

func TestWriteTransaction_shouldWriteATransaction(t *testing.T) {
	testDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	testFile := filepath.Join(testDir, "test.txt")

	o, err := NewFileOutput(testFile)
	if err != nil {
		log.Fatal(err)
	}
	err = o.WriteTransaction(transaction.Transaction{
		Hash:      "1",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})

	assert.NoError(t, err)
	output, err := ioutil.ReadFile(testFile)
	assert.NoError(t, err)
	assert.Equal(t, []byte("TxHash=1 Gas=10 FeePerGas=0.1 Signature=abc\n"), output)
}

func TestWriteTransaction_shouldWriteMultipleTransactions(t *testing.T) {
	testDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	testFile := filepath.Join(testDir, "test.txt")

	o, err := NewFileOutput(testFile)
	if err != nil {
		log.Fatal(err)
	}
	err = o.WriteTransaction(transaction.Transaction{
		Hash:      "1",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})
	assert.NoError(t, err)

	err = o.WriteTransaction(transaction.Transaction{
		Hash:      "2",
		Gas:       20,
		FeePerGas: 0.2,
		Signature: "def",
	})

	assert.NoError(t, err)
	output, err := ioutil.ReadFile(testFile)
	assert.NoError(t, err)
	assert.Equal(t, []byte("TxHash=1 Gas=10 FeePerGas=0.1 Signature=abc\n"+
		"TxHash=2 Gas=20 FeePerGas=0.2 Signature=def\n"), output)
}

func TestFeePerGas_should(t *testing.T) {
	assert.Equal(t, "0.8770012381849748", formatFeePerGas(float64(0.8770012381849748)))
}
