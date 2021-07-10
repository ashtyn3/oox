package network

import (
	"context"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
	"go.uber.org/zap"
)

var PORT uint16 = 5432

func (n *Node) Send(id noise.ID, s []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	n.Host.SendMessage(ctx, id.Address, Script{Contents: s})
	cancel()
}
func Handle(ctx noise.HandlerContext) error {
	if ctx.IsRequest() {
		return nil
	}

	obj, err := ctx.DecodeMessage()
	if err != nil {
		return nil
	}

	_, ok := obj.(Script)
	if !ok {
		return nil
	}
	return nil
}

func Cycle() {
	for {
	}
}
func NewFullNode(addr *net.IP, knownHosts []string, port ...uint16) {
	n := Node{}
	if len(port) > 0 {
		n = NewBasicNode(addr, &port[0], Handle)
	} else {
		n = NewBasicNode(addr, &PORT, Handle)
	}

	defer n.Host.Close()
	events := &kademlia.Events{OnPeerAdmitted: func(id noise.ID) {
		log.Info("New peer", zap.String("Id", id.Address))
	}}

	n.Bind(events, knownHosts)
	go Cycle()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Close stdin to kill the input goroutine.
	check(os.Stdin.Close())

	// Empty println.
	println()
}
