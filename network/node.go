package network

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ashtyn3/oox/utils"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
	"go.uber.org/zap"
)

type Node struct {
	Overlay *kademlia.Protocol
	Host    *noise.Node
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Script struct {
	Contents []byte
}

func (m Script) Marshal() []byte {
	return []byte(m.Contents)
}

func unmarshalScript(buf []byte) (Script, error) {
	return Script{Contents: buf}, nil
}

var log = utils.Logger()

// Returns unstrapped node
func NewBasicNode(host *net.IP, port *uint16, handler func(noise.HandlerContext) error) Node {
	// Create a new configured node.
	node, err := noise.NewNode(
		noise.WithNodeBindHost(*host),
		noise.WithNodeBindPort(*port),
	)
	check(err)

	// Register the chatMessage Go type to the node with an associated unmarshal function.
	node.RegisterMessage(Script{}, unmarshalScript)

	// Register a message handler to the node.
	node.Handle(handler)

	return Node{Host: node}
}

func (n *Node) Bind(events *kademlia.Events, ips []string) {
	overlay := kademlia.New(kademlia.WithProtocolEvents(*events))

	// Bind Kademlia to the node.
	n.Host.Bind(overlay.Protocol())

	n.Overlay = overlay

	n.strap(ips)
}

func (n *Node) strap(known []string) {
	// Have the node start listening for new peers.
	check(n.Host.Listen())

	// Ping nodes to initially bootstrap and discover peers from.
	bootstrap(n.Host, known...)

	// Attempt to discover peers if we are bootstrapped to any nodes.
	discover(n.Overlay)
}

// bootstrap pings and dials an array of network addresses which we may interact with and  discover peers from.
func bootstrap(node *noise.Node, addresses ...string) {
	for _, addr := range addresses {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_, err := node.Ping(ctx, addr)
		cancel()

		if err != nil {

			log.Error("Failed to ping node... Skipping", zap.String("Id", addr), zap.Error(err))
			continue
		}
	}
}

// discover uses Kademlia to discover new peers from nodes we already are aware of.
func discover(overlay *kademlia.Protocol) {
	ids := overlay.Discover()

	var str []string
	for _, id := range ids {
		str = append(str, fmt.Sprintf("%s(%s)", id.Address, id.ID.String()[:8]))
	}

	if len(ids) > 0 {
		log.Info("Found nodes", zap.Int("Number of nodes", len(ids)), zap.Any("Addresses", str))
	}
}
