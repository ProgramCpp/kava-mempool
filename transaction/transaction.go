package transaction

type Transaction struct {
	Hash      string
	Gas       int
	FeePerGas float64
	Signature string
}
