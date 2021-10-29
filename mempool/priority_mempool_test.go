package mempool

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/programcpp/kava-mempool/transaction"
)

func TestInsert_ShouldInsertNewTransactionIntoMemPool(t *testing.T) {
	testPool := NewMempool(1)

	testPool.Insert(transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	})

	assert.Equal(t, 1, testPool.Size())
	assert.Equal(t, transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	}, testPool.transactions[0])
}

func TestInsert_ShouldInsertMultipleTransactions(t *testing.T) {
	testPool := NewMempool(3)

	testPool.Insert(transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	})
	testPool.Insert(transaction.Transaction{
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	})

	assert.Equal(t, 2, testPool.Size())
	assert.Subset(t, []transaction.Transaction{{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	}, {
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	}}, testPool.transactions)
}
func TestInsert_ShouldDropLowestPrioritytransactionWhenCapacityIsFull(t *testing.T) {
	testPool := NewMempool(2)

	testPool.Insert(transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	})

	testPool.Insert(transaction.Transaction{
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	})

	testPool.Insert(transaction.Transaction{
		Hash:      "zzz",
		Gas:       300,
		FeePerGas: 0.1,
		Signature: "ghi",
	})

	assert.Equal(t, 2, testPool.Size())
	assert.Subset(t, []transaction.Transaction{{
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	}, {
		Hash:      "zzz",
		Gas:       300,
		FeePerGas: 0.1,
		Signature: "ghi",
	}}, testPool.transactions)

}

func TestRemove_ShouldRemoveLowestPriorityTransaction(t *testing.T) {
	testPool := NewMempool(2)

	testPool.Insert(transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	})
	testPool.Insert(transaction.Transaction{
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	})

	txn:= testPool.Remove()

	assert.Equal(t, transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	}, txn)

	assert.Equal(t, 1, testPool.Size())

	assert.Equal(t, transaction.Transaction{
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	}, testPool.transactions[0])
}

func TestInsert_ShouldInsertMultipleTransactionsPrioritizedByFee(t *testing.T) {
	testPool := NewMempool(3)

	testPool.Insert(transaction.Transaction{
		Hash:      "yyy",
		Gas:       40,
		FeePerGas: 0.1,
		Signature: "def",
	})
	testPool.Insert(transaction.Transaction{
		Hash:      "zzz",
		Gas:       50,
		FeePerGas: 0.1,
		Signature: "ghi",
	})
	testPool.Insert(transaction.Transaction{
		Hash:      "xxx",
		Gas:       30,
		FeePerGas: 0.1,
		Signature: "abc",
	})

	txn := testPool.Remove()
	assert.Equal(t, transaction.Transaction{
		Hash:      "xxx",
		Gas:       30, // lowest fee
		FeePerGas: 0.1,
		Signature: "abc",
	}, txn)

	txn = testPool.Remove()
	assert.Equal(t, transaction.Transaction{
		Hash:      "yyy",
		Gas:       40, // second lowest fee
		FeePerGas: 0.1,
		Signature: "def",
	}, txn)

	txn = testPool.Remove()
	assert.Equal(t, transaction.Transaction{
		Hash:      "zzz",
		Gas:       50, // highest fee
		FeePerGas: 0.1,
		Signature: "ghi",
	}, txn)
}