package internal

import (
	"testing"

	"github.com/programcpp/kava-mempool/transaction"
	"github.com/stretchr/testify/assert"
)

func TestLen_ShouldReturnTheLengthOfHeap(t *testing.T) {
	h := TransactionsHeap{}

	assert.Zero(t, h.Len())

	h = append(h, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})

	assert.Equal(t, 1, h.Len())

	h = append(h, transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "def",
	})

	assert.Equal(t, 2, h.Len())
}

func TestLess_ShouldReturnTrueWhenFirstElementHasLowerGasValue(t *testing.T) {
	h := TransactionsHeap{}

	h = append(h, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})
	h = append(h, transaction.Transaction{
		Hash:      "yyy",
		Gas:       20,
		FeePerGas: 0.1,
		Signature: "def",
	})

	assert.True(t, h.Less(0, 1))
}

func TestLess_ShouldReturnTrueWhenFirstElementHasLowerFeeValue(t *testing.T) {
	h := TransactionsHeap{}

	h = append(h, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})
	h = append(h, transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.2,
		Signature: "def",
	})

	assert.True(t, h.Less(0, 1))
}

func TestSwap(t *testing.T) {
	h := TransactionsHeap{}

	h = append(h, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})
	h = append(h, transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.2,
		Signature: "def",
	})

	h.Swap(0, 1)

	assert.Equal(t, transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.2,
		Signature: "def",
	}, h[0])

	assert.Equal(t, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	}, h[1])
}

func TestPush(t *testing.T) {
	h := TransactionsHeap{}
	h.Push(transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	})

	assert.Equal(t, 1, len(h))
	assert.Equal(t, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	}, h[0])

	h.Push(transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.2,
		Signature: "def",
	})

	assert.Equal(t, 2, len(h))
	assert.Equal(t, transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.2,
		Signature: "def",
	}, h[1])
}

func TestPop(t *testing.T) {
	h := TransactionsHeap{
		{
			Hash:      "xxx",
			Gas:       10,
			FeePerGas: 0.1,
			Signature: "abc",
		}, {
			Hash:      "yyy",
			Gas:       10,
			FeePerGas: 0.2,
			Signature: "def",
		},
		{
			Hash:      "zzz",
			Gas:       10,
			FeePerGas: 0.3,
			Signature: "ghi",
		},
	}

	txn := h.Pop()
	assert.Equal(t, 2, h.Len())
	assert.Equal(t, transaction.Transaction{
		Hash:      "zzz",
		Gas:       10,
		FeePerGas: 0.3,
		Signature: "ghi",
	}, txn)

	txn = h.Pop()
	assert.Equal(t, 1, h.Len())
	assert.Equal(t, transaction.Transaction{
		Hash:      "yyy",
		Gas:       10,
		FeePerGas: 0.2,
		Signature: "def",
	}, txn)

	txn = h.Pop()
	assert.Equal(t, 0, h.Len())
	assert.Equal(t, transaction.Transaction{
		Hash:      "xxx",
		Gas:       10,
		FeePerGas: 0.1,
		Signature: "abc",
	}, txn)
}
