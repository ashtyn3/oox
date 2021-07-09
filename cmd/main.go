package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ashtyn3/oox/core/txn"
	"github.com/ashtyn3/oox/wallet"
)

func main() {

	// KYNY7YGboLEAF6PUUtw8ejFMtpVbDobHH

	p, _ := wallet.Get("wow", "Hello1", "KYNY7YGboLEAF6PUUtw8ejFMtpVbDobHH")
	t := txn.Transaction{Out: append([][]byte{}, p.Public), Value: 20}
	t.Hash()

	r, s, _ := ecdsa.Sign(rand.Reader, p.Private, []byte(t.Id))

	sig := append(r.Bytes(), s.Bytes()...)
	t.Sig = hex.EncodeToString(sig)

	x := &big.Int{}
	x.SetBytes(t.Out[0][:len(t.Out[0])/2])

	y := &big.Int{}
	y.SetBytes(t.Out[0][len(t.Out[0])/2:])

	R := &big.Int{}
	R.SetBytes(sig[:len(sig)/2])

	S := &big.Int{}
	S.SetBytes(sig[len(sig)/2:])

	curve := elliptic.P521()
	pub := ecdsa.PublicKey{curve, x, y}

	v := ecdsa.Verify(&pub, []byte(t.Id), R, S)

	fmt.Println(v)
}
