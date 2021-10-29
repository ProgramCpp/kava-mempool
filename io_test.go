package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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

	file, err := NewFileIO(filepath.Join(testDir, inputFile))
	if err != nil {
		log.Fatal(err)
	}
	txn, err := file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, "", txn.hash)
	assert.Equal(t, 0, txn.gas)
	assert.Equal(t, float32(0), txn.feePerGas)
	assert.Equal(t, "", txn.signature)
}

func TestReadTransaction_ShouldReturnTheTransaction(t *testing.T) {
	inputFile := "testFile"
	testDir := testSetup(inputFile, []byte("TxHash=0x54030E30503453949230403 Gas=300000 FeePerGas=0.001 Signature=0x54030E30503453949230403"))
	defer os.RemoveAll(testDir)

	file, err := NewFileIO(filepath.Join(testDir, inputFile))
	if err != nil {
		log.Fatal(err)
	}
	txn, err := file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, "0x54030E30503453949230403", txn.hash)
	assert.Equal(t, 300000, txn.gas)
	assert.Equal(t, float32(.001), txn.feePerGas)
	assert.Equal(t, "0x54030E30503453949230403", txn.signature)
}

func TestReadTransaction_ShouldReturnOneTransactionAtATime(t *testing.T) {
	inputFile := "testFile"
	testDir := testSetup(inputFile, []byte(
		"TxHash=1 Gas=3 FeePerGas=0.001 Signature=abc\n"+
			"TxHash=2 Gas=4 FeePerGas=0.001 Signature=def"))
	defer os.RemoveAll(testDir)

	file, err := NewFileIO(filepath.Join(testDir, inputFile))
	if err != nil {
		log.Fatal(err)
	}
	txn, err := file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, transaction{
		hash:      "1",
		gas:       3,
		feePerGas: 0.001,
		signature: "abc",
	}, txn)

	txn, err = file.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, transaction{
		hash:      "2",
		gas:       4,
		feePerGas: 0.001,
		signature: "def",
	}, txn)
}
