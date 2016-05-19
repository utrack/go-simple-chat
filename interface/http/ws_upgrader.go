package ifaceHttp

import (
	"github.com/gorilla/websocket"
	"github.com/utrack/go-simple-chat/client/ws"
	"github.com/utrack/go-simple-chat/hub"
	"github.com/utrack/go-simple-chat/message"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWs handles websocket requests from the peer.
func ServeWs(h hub.ClientAcceptor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := clientWs.NewClient(ws)
		err = h.RegisterClient(client, r.URL.Query().Get(`name`))
		if err != nil {
			client.Send(message.One{Type: message.EventError, Text: err.Error()})
			client.Disconnect()
		}
	}
}
