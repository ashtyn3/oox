package crypto

import (
	"fmt"

	"github.com/mr-tron/base58"
)

// Base58 encoder function
func B58Enc(in []byte) []byte {
	b := base58.Encode(in)
	return []byte(b)
}

func B58Dec(in []byte) []byte {
	b, err := base58.Decode(string(in))

	if err != nil {
		fmt.Println(err)
	}
	return b
}
