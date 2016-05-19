package hub

import (
	"github.com/utrack/go-simple-chat/client"
)

// Hub routes messages between clients and
// controls the clients connected to it.
type Hub interface {
	// RegisterClient adds the client with given nickname
	// to the hub and starts processing its messages.
	RegisterClient(client.Client, string) error
	// Run starts the hub's messaging pump.
	Run()
}
