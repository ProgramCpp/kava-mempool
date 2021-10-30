# Priority Mempool
A priority mempool that orders transactions by fee. this mempool is useful to prioritize transactions of higher fee and hence deliver higher transction commits.

this priority mempool,
- reads transactions from a file, stores them ordered by __lowest transaction fee* to highest__
- drops lowest priority trnsaction once the mempool capacity is full

\* transaction fee = Gas * FeePerGas

example transaction format,
```
TxHash=0x54030E30503453949230403 Gas=300000 FeePerGas=0.001 Signature=0x54030E30503453949230403
```
example of mempool with capacity 3:
input:
```
TxHash=4 Gas=6 FeePerGas=0.001 Signature=jkl
TxHash=2 Gas=4 FeePerGas=0.001 Signature=def
TxHash=1 Gas=3 FeePerGas=0.001 Signature=abc
TxHash=3 Gas=5 FeePerGas=0.001 Signature=ghi

```
output:
```
TxHash=2 Gas=4 FeePerGas=0.001 Signature=def
TxHash=3 Gas=5 FeePerGas=0.001 Signature=ghi
TxHash=4 Gas=6 FeePerGas=0.001 Signature=jkl
```
In this example, 
- transactions are orders by fee
- the transaction with TxHash `1` gets dropped with the lowest fee

## Pre-requisites
- golang version 1.17 or higher


## Build Steps
run 
```
go build 
```


## Unit Tests
run
```
go test ./...
```


## Usage
run 
```
./kava-mempool
```
the executable by default,
1. builds mempool of capacity 5000
2. reads transactions from `./transactions.txt` and writes prioritized transactions to `./prioritized-transactions.txt`


## ADR
### [10/28/2021] 
#### Summary: 
file IO is stateful. i.e, each instance of IO reads/writes only one transaction at a time and tracks the point till which the transaction is read/written and resumes reading/ writing transactions from that point. 
#### Background
The transaction processing is time sensitive and needs to be processsed as soon as it is made available
#### Decision: 
do not read the whole file at once
#### Pros: 
- This helps adding transactions into mempool as it arrives at the node
- This generic interface implementation can be swapped for other source of transactions. ex: network
#### Cos: 
- disk io is not efficient
- File object is stateful. The outcome of the input utility is not idempotent
- Its Cleanup is decoupled from the initialization code


## Future Work
 - add Makefile for build and test steps
 - process HIGHEST priority transaction first. i.e, order transactions by highest transaction fee to lowest transaction fee
 - create sub packages for different concerns in order to improve readability and maintainability. ex: processor, io
 - move private members of the package to internal package
 - accept mempool capacity, input and output files as command line arguments


## Reference
 [https://www.blocknative.com/blog/mempool-intro](https://www.blocknative.com/blog/mempool-intro)

