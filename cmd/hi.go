package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ashtyn3/oox/network"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
	"github.com/spf13/pflag"
)

var (
	hostFlag = pflag.IPP("host", "h", nil, "binding host")
	portFlag = pflag.Uint16P("port", "p", 0, "binding port")
	sendFlag = pflag.Bool("send", false, "Sets mode to txn sending")
)

// check panics if err is not nil.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// printedLength is the total prefix length of a public key associated to a chat users ID.

func main() {
	// Parse flags/options.
	pflag.Parse()

	n := network.NewBasicNode(hostFlag, portFlag)
	overlay := n.Overlay
	node := n.Host

	defer node.Close()

	if *sendFlag == true {
		ids := overlay.Table().Peers()

		for _, id := range ids {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			node.SendMessage(ctx, id.Address, network.Script{Contents: []byte("Hello world")})
			cancel()
		}
		check(os.Stdin.Close())
		return
	}
	// Accept chat message inputs and handle chat commands in a separate goroutine.
	go func() {
		for {
		}
	}()

	// Wait until Ctrl+C or a termination call is done.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Close stdin to kill the input goroutine.
	check(os.Stdin.Close())

	// Empty println.
	println()
}

// Send sends data to connected nodes
func send(node *noise.Node, overlay *kademlia.Protocol, line []byte) {
	for _, id := range overlay.Table().Peers() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		err := node.SendMessage(ctx, id.Address, network.Script{Contents: line})
		cancel()

		if err != nil {
			fmt.Printf("Failed to send message to %s(%s). Skipping... [error: %s]\n",
				id.Address,
				id.ID.String()[:printedLength],
				err,
			)
			continue
		}
	}
}
