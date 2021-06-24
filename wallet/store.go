package wallet

import (
	"bytes"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ashtyn3/oox/utils/crypto"
)

type Options struct {
	Path     string
	Password string
}

func InitWallet(options *Options, pairs []KeyPair) error {
	if options.Password == "" {
		return errors.New("Password not provided")
	}

	if options.Path == "" {
		options.Path = "./"
	}

	var content bytes.Buffer

	gob.Register(elliptic.P256())

	enc := gob.NewEncoder(&content)
	enc.Encode(pairs)
	sha := sha256.Sum256([]byte(options.Password))
	e, err := crypto.Encrypt(sha[:], content.Bytes())

	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile(options.Path+fmt.Sprintf("%x", pairs[0].Public)[0:11], e, 0644)
	return nil
}
