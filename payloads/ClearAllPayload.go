package payloads

// NewClearAllPayload creates a new Cleal All Payload
func NewClearAllPayload() Payload {
	return Payload {
		Type: "clear_all",
	}
}
