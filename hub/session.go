package hub

import (
	"github.com/utrack/go-simple-chat/client"
	"github.com/utrack/go-simple-chat/message"
)

// session is a middleware between a client
// and the hub.
//
// session forwards messages and discon notifications
// from a client to the hub's main chan.
type session struct {
	// underlying client
	c client.Client

	// client's name
	name string

	// hub's main incoming messages' channel
	chatChannelChan chan<- message.One

	disconChan chan<- disconMsg
}

// newSession returns a forged session for given
// client.
// args: client, its name, channel to dump incoming messages to.
func newSession(c client.Client, name string, ch chan<- message.One, dch chan<- disconMsg) *session {
	return &session{
		c:               c,
		name:            name,
		chatChannelChan: ch,
		disconChan:      dch,
	}
}

func (s *session) runPump() {
	incoming := s.c.MsgChan()
	discon := s.c.DisconChan()
	for {
		select {
		case msg := <-incoming:
			// Filter out malicious EventTypes
			switch msg.Type {
			case message.EventMessage:
			case message.EventPresenceState:
			default:
				continue
			}

			msg.From = s.name
			s.chatChannelChan <- msg
		case disconReason := <-discon:
			// Send Leaving... message
			s.chatChannelChan <- message.One{
				Type: message.EventLeave,
				From: s.name,
				Text: disconReason.Error(),
			}
			s.disconChan <- disconMsg{name: s.name, reason: disconReason.Error()}
			break
		}
	}
}

func (s *session) send(m message.One) error {
	return s.c.Send(m)
}
