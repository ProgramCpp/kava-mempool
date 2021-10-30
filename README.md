# Priority Mempool
A priority mempool that orders transactions by fee. this mempool is useful to prioritize transactions of higher fee and hence deliver higher transction commits.


## Pre-requisites
- golang version 1.17 or higher


## Build Steps

## Unit Tests


## Run Steps


## ADR
### [10/28/2021] 
#### Summary: 
file input is stateful. i.e, each instance of IO reads only one transaction at a time and tracks the point till which the transaction is read and resumes reading transactions from that point. 
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
 - process HIGHEST priority transaction first
 - create sub packages for different concerns in order to improve readability and maintainability. ex: processor, io
 - move private members of the package to internal package
 - accept input and output files as command line arguments


## Reference
 [https://www.blocknative.com/blog/mempool-intro](https://www.blocknative.com/blog/mempool-intro)

