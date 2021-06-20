package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"

	"github.com/ashtyn3/oox/core/txn"
	"github.com/ashtyn3/oox/utils"
)

type Data struct {
	B float64
	G float64
}

type Block struct {
	Txns      []txn.Transaction
	Hash      string
	Prev      []byte
	validator string
	BlockData Data
}

// Finds hash and assigns it to Hash feild in block.
func (b *Block) InitHash() {
	var encoder bytes.Buffer

	gob.Register(new(Block))
	encode := gob.NewEncoder(&encoder)
	encode.Encode(b)

	sha := sha256.Sum256(encoder.Bytes())

	b.Hash = hex.EncodeToString(sha[:])
}

// Calculates block score and assigns its value to B feild in Data struct.
func (b *Block) BlockScore() {
	b.BlockData.B = utils.Stdev(b.Txns) + float64(utils.Max(b.Txns))
}

// Calculates block wide gas price and assigns its value to G feild in Data struct.
func (b *Block) BWGS() {
	b.BlockData.G = b.BlockData.B / 10
}

func CreateBlock(txns []txn.Transaction, prevHash []byte) {
	D := Data{
		B: 0,
		G: 0,
	}
	block := &Block{Txns: txns, Prev: prevHash, BlockData: D}
	block.InitHash()

}
