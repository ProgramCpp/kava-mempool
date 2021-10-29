package internal

import "github.com/programcpp/kava-mempool/transaction"

// implements heap.Interface to keep the transactions ordered by fee.
// higher fee has higher priority
type TransactionsHeap []transaction.Transaction

func (t TransactionsHeap) Len() int { return len(t) }
func (t TransactionsHeap) Less(i, j int) bool {
	return t[i].FeePerGas*float32(t[i].Gas) < t[j].FeePerGas*float32(t[j].Gas)
}
func (t TransactionsHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *TransactionsHeap) Push(txn interface{}) {
	*t = append(*t, txn.(transaction.Transaction))
}

func (t *TransactionsHeap) Pop() interface{} {
	n := len(*t)
	txn := (*t)[n-1]
	*t = (*t)[0 : n-1]
	return txn
}

