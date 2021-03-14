package payloads

// NewShowAppPayload creates a new Show App Payload
func NewShowAppPayload() Payload {
	return Payload {
		Type: "show_app",
	}
}
