package message

// EventType marks an event's type.
type EventType uint64

const (
	// EventUnknown is the default EventType.
	EventUnknown = iota
	// EventJoin is sent when some user joins the channel.
	EventJoin
	// EventLeave is sent when some user leaves the channel.
	EventLeave
	// EventPresenceState is sent when presence state of a user changes.
	EventPresenceState
	// EventMessage is sent when a message was received.
	EventMessage
	// EventError is sent when an error had occurred. Normally it's sent only
	// during the connection establishment phase.
	EventError
)
