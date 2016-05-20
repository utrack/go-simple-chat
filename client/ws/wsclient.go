//Package clientWs provides Websocket transport for the Client.
package clientWs

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/pquerna/ffjson/ffjson"
	clientIface "github.com/utrack/go-simple-chat/client"
	"github.com/utrack/go-simple-chat/message"
	"sync"
	"time"
)

type client struct {
	ws *websocket.Conn

	disconChan chan error

	// messages that are awaiting their turn to be forwarded to the client.
	messagesToSend       chan message.One
	messagesToSendClosed bool
	// protects messagesToSend/messagesToSendClosed
	messagesToSendMu sync.Mutex

	// messages rcvd from the client.
	messagesRcvd chan message.One
}

// NewClient returns a client.Client chat connection for
// provided WebSocket connection.
func NewClient(ws *websocket.Conn) clientIface.Client {
	ret := &client{
		ws:             ws,
		disconChan:     make(chan error, 2),
		messagesToSend: make(chan message.One, 15),
		messagesRcvd:   make(chan message.One, 15),
	}
	go ret.writePump()
	go ret.readPump()
	return ret
}

func (c *client) DisconChan() <-chan error {
	return c.disconChan
}

func (c *client) MsgChan() <-chan message.One {
	return c.messagesRcvd
}

func (c *client) Send(m message.One) error {
	c.messagesToSendMu.Lock()
	defer c.messagesToSendMu.Unlock()

	if c.messagesToSendClosed {
		return nil
	}
	c.messagesToSend <- m
	return nil
}

func (c *client) Disconnect() {
	c.discon(errors.New("Kicked"))
}

func (c *client) discon(reason error) {
	c.messagesToSendMu.Lock()
	defer c.messagesToSendMu.Unlock()

	if c.messagesToSendClosed {
		return
	}

	c.messagesToSendClosed = true
	close(c.messagesToSend)

	c.disconChan <- reason
}

// errDisconGoingAway is returned to the disconChan if the client had gracefully closed the channel.
var errDisconGoingAway = errors.New("Leaving")

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// readPump pumps messages from the websocket connection.
func (c *client) readPump() {
	errDiscon := errDisconGoingAway

	// Forward the discon info on disconnect
	defer func() {
		c.discon(errDiscon)
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	var msg message.One
	for {
		_, dataPkg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				errDiscon = err
			}
			break
		}

		err = ffjson.Unmarshal(dataPkg, &msg)
		if err != nil {
			break
		}
		c.messagesRcvd <- msg
	}
}

func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	var err = errDisconGoingAway
	// Forward the discon info on disconnect
	defer func() {
		c.discon(err)
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.messagesToSend:
			// Input chan was closed, shutdown the sock
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			// Marshal JSON and send
			buf, _ := ffjson.Marshal(message)
			if err = c.write(websocket.TextMessage, buf); err != nil {
				return
			}
			ffjson.Pool(buf)

		case <-ticker.C:
			if err = c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// write writes a message with the given message type and payload.
func (c *client) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}
