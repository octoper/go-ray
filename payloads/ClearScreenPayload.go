package payloads

// NewClearScreenPayload creates a new Clear Screen Payload
func NewClearScreenPayload() Payload {
	return Payload {
		Type: "new_screen",
		Content: "",
	}
}
