package transaction

type Transaction struct {
	Hash string
	Gas int
	FeePerGas float32
	Signature string
}