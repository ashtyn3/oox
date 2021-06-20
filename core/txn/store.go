package txn

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

func (tx *Transaction) Hash() {
	var encoded bytes.Buffer

	encode := gob.NewEncoder(&encoded)
	encode.Encode(tx)

	hash := sha256.Sum256(encoded.Bytes())
	tx.Id = hex.EncodeToString(hash[:])
}
