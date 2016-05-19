package hub

import (
	"github.com/utrack/go-simple-chat/client"
	"github.com/utrack/go-simple-chat/message"
	"sync"
)

// Hub routes messages between clients and
// controls the clients connected to it.
type Hub struct {
	// incomingMsgs contains all the messages
	// rcvd from the clients. Pre-processing queue.
	incomingMsgs chan message.One

	// incomingDiscons contains incoming discon
	// notifications.
	incomingDiscons chan disconMsg

	// sessions is a map of nickname->session. Contains
	// every connected session.
	sessions map[string]*session

	// sessionsMu protects the sessions' map.
	sessionsMu sync.RWMutex
}

// NewHub initiates and returns the Hub.
// Execute Hub.Run() to run the processing pumps.
func NewHub() *Hub {
	return &Hub{
		incomingMsgs:    make(chan message.One, 45),
		incomingDiscons: make(chan disconMsg, 45),

		sessions:   make(map[string]*session),
		sessionsMu: sync.RWMutex{},
	}
}

// RegisterClient adds the client to the Hub.
func (h *Hub) RegisterClient(c client.Client, name string) error {
	if h.clientExists(name) {
		return ErrNickCollision
	}

	sess := newSession(c, name, h.incomingMsgs, h.incomingDiscons)
	h.addSession(sess)
	return nil
}

// Run starts the message processing pump which accepts
// messages from the clients and routes them around.
func (h *Hub) Run() {
	go h.pump()
}

// pump processes incoming messages and discon notifications.
func (h *Hub) pump() {
	for {
		select {
		case msg := <-h.incomingMsgs:
			go h.sendMsg(msg)
		case discon := <-h.incomingDiscons:
			h.removeSession(discon.name)
			// TODO log debug
		}
	}
}

// removeSession removes the session from sessions' dict.
func (h *Hub) removeSession(key string) {
	h.sessionsMu.Lock()
	defer h.sessionsMu.Unlock()
	delete(h.sessions, key)
}

// addSession inserts the session to the sessions' map.
func (h *Hub) addSession(s *session) {
	h.sessionsMu.Lock()
	defer h.sessionsMu.Unlock()

	h.sessions[s.name] = s
	// TODO log debug
}

// sendMsg distributes the message to every connected session.
// It's safe to run sendMsg concurrently.
func (h *Hub) sendMsg(m message.One) {
	h.sessionsMu.RLock()
	defer h.sessionsMu.RUnlock()

	var err error
	for _, sess := range h.sessions {
		err = sess.send(m)
		err.Error()
		// TODO log debug
		// Sessions and clients handle error thresholds themselves,
		// so hub shouldn't discon the session on its own.
	}
}

// clientExists returns true if the client with given name was found.
func (h *Hub) clientExists(name string) bool {
	h.sessionsMu.RLock()
	defer h.sessionsMu.RUnlock()
	_, ok := h.sessions[name]
	return ok
}
