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


func testSetup(file string, content []byte) string{
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
func TestReadTransaction_fromEmptyFileShouldReturnEmptyTransaction(t *testing.T) {
	inputFile := "testFile"
	testDir := testSetup(inputFile, []byte(""))
	defer os.RemoveAll(testDir)

	txn, err := fileIO{inputFilePath: filepath.Join(testDir, inputFile)}.readTransaction()
	assert.NoError(t, err)
	assert.Equal(t, "", txn.hash)
	assert.Equal(t, 0, txn.gas)
	assert.Equal(t, float32(0), txn.feePerGas)
	assert.Equal(t, "", txn.signature)
}
