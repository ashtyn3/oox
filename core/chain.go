package core

import "github.com/syndtr/goleveldb/leveldb"

func NewChain(genesis *Block, id string) {
	db, _ := leveldb.OpenFile(id+".OOX_CHAIN", nil)

	db.Put([]byte(genesis.Hash), genesis.Bytes(), nil)
}
