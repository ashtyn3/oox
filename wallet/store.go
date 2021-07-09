package wallet

import (
	"bytes"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"log"
	"os"

	"github.com/ashtyn3/oox/utils/crypto"
)

type Options struct {
	Path     string
	Password string
}

type Wallet map[string]*KeyPair

func (w *Wallet) InitWallet(name string, options *Options) error {
	if options.Password == "" {
		return errors.New("Password not provided")
	}

	if options.Path == "" {
		options.Path = "./"
	}

	var content bytes.Buffer

	//gob.Register(elliptic.P256())
	gob.Register(&elliptic.CurveParams{})

	enc := gob.NewEncoder(&content)
	err := enc.Encode(w)

	if err != nil {
		log.Fatalln(err)
	}
	sha := sha256.Sum256([]byte(options.Password))
	e, err := crypto.Encrypt(sha[:], content.Bytes())

	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(options.Path+name+".OOX_WALLET", e, 0644)

	return nil
}

func Add(name, pwd string, pair *KeyPair) {
	m := GetPairs(name, pwd)

	addr := pair.Addr()
	m[addr] = pair

	m.InitWallet(name, &Options{Password: pwd})
}

func Get(name, pwd string, addr string) (*KeyPair, error) {
	m := GetPairs(name, pwd)

	if m[addr] == nil {
		return nil, errors.New(addr + " does not exist in wallet")
	}

	return m[addr], nil
}

func GetPairs(name, pwd string) Wallet {
	sha := sha256.Sum256([]byte(pwd))
	EContent, _ := os.ReadFile(name + ".OOX_WALLET")
	DContent, err := crypto.Decrypt(sha[:], EContent)

	if err != nil {
		log.Fatalln(err)
	}

	var kpArray Wallet

	//gob.Register(elliptic.P256())
	gob.Register(&elliptic.CurveParams{})

	dec := gob.NewDecoder(bytes.NewReader(DContent))
	err = dec.Decode(&kpArray)

	if err != nil {
		log.Fatalln(err)
	}

	return kpArray
}
