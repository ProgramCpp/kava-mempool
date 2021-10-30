package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/programcpp/kava-mempool/transaction"
)

func testSetup(file string, content []byte) string {
	testDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	tmpFile := filepath.Join(testDir, file)
	fmt.Println(tmpFile)
	err = ioutil.WriteFile(tmpFile, content, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return testDir
}
func TestReadTransaction_ShouldReturnEmptyTransactionfromEmptyFile(t *testing.T) {
	inputFile := "testFile"
	testDir := testSetup(inputFile, []byte(""))
	defer os.RemoveAll(testDir)

	file, err := NewFileInput(filepath.Join(testDir, inputFile))
	if err != nil {
		log.Fatal(err)
	}
	txn, err := file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, "", txn.Hash)
	assert.Equal(t, 0, txn.Gas)
	assert.Equal(t, float64(0), txn.FeePerGas)
	assert.Equal(t, "", txn.Signature)
}

func TestReadTransaction_ShouldReturnTheTransaction(t *testing.T) {
	inputFile := "testFile"
	testDir := testSetup(inputFile, []byte("TxHash=0x54030E30503453949230403 Gas=300000 FeePerGas=0.001 Signature=0x54030E30503453949230403"))
	defer os.RemoveAll(testDir)

	file, err := NewFileInput(filepath.Join(testDir, inputFile))
	if err != nil {
		log.Fatal(err)
	}
	txn, err := file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, "0x54030E30503453949230403", txn.Hash)
	assert.Equal(t, 300000, txn.Gas)
	assert.Equal(t, float64(.001), txn.FeePerGas)
	assert.Equal(t, "0x54030E30503453949230403", txn.Signature)
}

func TestReadTransaction_ShouldReturnOneTransactionAtATime(t *testing.T) {
	inputFile := "testFile"
	testDir := testSetup(inputFile, []byte(
		"TxHash=1 Gas=3 FeePerGas=0.001 Signature=abc\n"+
			"TxHash=2 Gas=4 FeePerGas=0.001 Signature=def"))
	defer os.RemoveAll(testDir)

	file, err := NewFileInput(filepath.Join(testDir, inputFile))
	if err != nil {
		log.Fatal(err)
	}
	txn, err := file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, transaction.Transaction{
		Hash:      "1",
		Gas:       3,
		FeePerGas: 0.001,
		Signature: "abc",
	}, txn)

	txn, err = file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, transaction.Transaction{
		Hash:      "2",
		Gas:       4,
		FeePerGas: 0.001,
		Signature: "def",
	}, txn)
}
