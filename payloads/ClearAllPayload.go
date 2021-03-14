package payloads

// NewClearAllPayload creates a new Clear All Payload
func NewClearAllPayload() Payload {
	return Payload {
		Type: "clear_all",
	}
}
