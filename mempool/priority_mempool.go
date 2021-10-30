package mempool

import (
	"container/heap"

	"github.com/programcpp/kava-mempool/mempool/internal"
	"github.com/programcpp/kava-mempool/transaction"
	log "github.com/sirupsen/logrus"
)


// A priority mem pool that orders the transactions by fee
// a transaction with higher fee has higher priority
// insertion takes O(log N)
type PriorityMempool struct {
	transactions internal.TransactionsHeap
	// maximum transactions the mem pool can hold
	capacity int
}

func NewMempool(capacity int) PriorityMempool {
	return PriorityMempool{capacity: capacity}
}

// returns the number of transactiosn present in the mem pool
func (p *PriorityMempool) Size() int {
	return len(p.transactions)
}

// inserts transaction into mem Pool
// drops the lowest priority transaction when the capacity is full
func (p *PriorityMempool) Insert(t transaction.Transaction) {
	if len(p.transactions) == p.capacity {
		t := heap.Pop(&p.transactions)
		log.Info("dropping lowest priority transaction: ", t)
	}
	heap.Push(&p.transactions, t)
}

// removes the lowest priority transaction
func (p *PriorityMempool) Remove() transaction.Transaction {
	if len(p.transactions) == 0 {
		return transaction.Transaction{}
	}
	t := heap.Pop(&p.transactions)
	return t.(transaction.Transaction)
}
