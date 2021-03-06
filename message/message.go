//Package message provides data structures that get transferred between the client and server.
package message

// One is the data package (message) about an event
// that gets sent to the client.
type One struct {
	// Type is the message's event type.
	Type EventType `json:"event_type"`
	// IsMuted is true if the event shouldn't
	// show up in the channel's logs.
	IsMuted bool `json:"is_muted"`
	// From is the message's source nickname.
	From string `json:"from"`
	// Text is the event's aux text.
	Text string `json:"text"`
}
