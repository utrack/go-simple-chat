package hub

// disconMsg is sent by session to the hub;
// notifies that some client has gone away.
type disconMsg struct {
	// client's name
	name string
	// discon reason
	reason string
}
