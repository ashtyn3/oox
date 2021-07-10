package network

import (
	"context"
	"fmt"
	"time"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)

func send(node *noise.Node, overlay *kademlia.Protocol, line []byte) {
	for _, id := range overlay.Table().Peers() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		err := node.SendMessage(ctx, id.Address, Script{Contents: line})
		cancel()

		if err != nil {
			fmt.Printf("Failed to send message to %s. Skipping... [error: %s]\n",
				id.Address,
				err,
			)
			continue
		}
	}
}
