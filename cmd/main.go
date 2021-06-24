package main

import (
	"fmt"

	"github.com/ashtyn3/oox/wallet"
)

func main() {
	k := wallet.NewKeyPair()
	fmt.Printf("Address: %s\n", k.Addr())

	wallet.InitWallet(&wallet.Options{Path: "./", Password: "Hello1"}, []wallet.KeyPair{k})

}
