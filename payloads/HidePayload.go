package payloads

// NewHidePayload creates a new Hide Payload
func NewHidePayload() Payload {
	return Payload {
		Type: "hide",
	}
}
