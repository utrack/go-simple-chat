//Package client provides abstraction over the client-server transport.
package client

import (
	"github.com/utrack/go-simple-chat/message"
)

// Client is a common interface to clients' incoming connections.
type Client interface {
	// Send sends the message packet to the client. Returns err on error.
	Send(message.One) error
	// DisconChan returns a client's disconnect notification channel,
	// which receives single value (error) at most. Error contains the
	// discon's description which describes why the client has gone away.
	DisconChan() <-chan error

	// MsgChan returns channel that receives incoming messages from the
	// client.
	MsgChan() <-chan message.One

	// Disconnect forcefully closes the client's connection.
	Disconnect()
}
