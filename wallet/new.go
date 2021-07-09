package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/ashtyn3/oox/utils/crypto"
	"golang.org/x/crypto/ripemd160"
)

type KeyPair struct {
	Private *ecdsa.PrivateKey
	Public  []byte
}

const (
	Checksum = 4
)

func NewKeyPair() *KeyPair {
	c := elliptic.P521()
	pk, _ := ecdsa.GenerateKey(c, rand.Reader)
	pub := append(pk.PublicKey.X.Bytes(), pk.PublicKey.Y.Bytes()...)

	return &KeyPair{Private: pk, Public: pub}
}

func keyHash(key []byte) []byte {
	pubHash := sha256.Sum256(key)

	r := ripemd160.New()
	r.Write(pubHash[:])

	return r.Sum(nil)
}

func checksum(hash []byte) []byte {
	f := sha256.Sum256(hash)
	s := sha256.Sum256(f[:])

	return s[:Checksum]
}

func (kp *KeyPair) Addr() string {
	hash := keyHash(kp.Public)
	checkHash := checksum(hash)
	fullHash := append(hash, checkHash...)

	addr := crypto.B58Enc(fullHash)

	return string(addr)
}

func Partial(rawKey []byte) string {
	hash := keyHash(rawKey)
	checkHash := checksum(hash)
	fullHash := append(hash, checkHash...)

	addr := crypto.B58Enc(fullHash)

	return string(addr)
}
