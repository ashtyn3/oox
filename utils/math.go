package utils

import (
	"math"

	"github.com/ashtyn3/oox/core/txn"
)

// Finds standard deviation of array of transaction's values.
func Stdev(txs []txn.Transaction) float64 {
	sumSq := 0
	sum := 0

	for _, t := range txs {
		sum += t.Value * t.Value
		sumSq += t.Value
	}

	meanSq := sumSq / len(txs)
	mean := sum / len(txs)

	stdev := math.Sqrt(float64(meanSq) - float64(mean*mean))

	return stdev
}

func Max(txs []txn.Transaction) int {
	max := txs[0].Value
	for _, t := range txs {
		if t.Value > max {
			max = t.Value
		}
	}
	return max
}
