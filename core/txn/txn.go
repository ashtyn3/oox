package txn

type Transaction struct {
	Out   [][]byte
	Value int
	Id    string
	Sig   string
}
