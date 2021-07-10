package main

import (
	"github.com/ashtyn3/oox/network"
	"github.com/spf13/pflag"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	p := pflag.UintP("port", "p", 5432, "Port number")
	ip := pflag.IPP("address", "a", nil, "Port number")

	pflag.Parse()
	if p != nil {
		network.NewFullNode(ip, pflag.Args(), uint16(*p))
	} else {
		network.NewFullNode(ip, pflag.Args())
	}
}
