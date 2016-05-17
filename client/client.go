package client

import (
	"github.com/utrack/go-simple-chat/message"
)

type Client interface {
	// Send sends the message packet to the client. Returns err on error.
	Send(message.One) error
	// DisconChan returns a client's disconnection notification channel,
	// which receives single value (error) at most. Error contains the
	// discon's description which describes why the client has gone away.
	DisconChan() <-chan error
}
