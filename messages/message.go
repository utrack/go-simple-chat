package messages

// Message is the data package about an event
// that gets sent to the client.
type Message struct {
	// Type is the message's event type.
	Type EventType `json:"event_type"`
	// IsMuted is true if the event shouldn't
	// show up in the channel's logs.
	IsMuted bool `json:"is_muted"`
	// Text is the event's aux text.
	Text string `json:"text"`
}
