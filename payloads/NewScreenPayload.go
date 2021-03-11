package payloads

// NewNewScreenPayload creates a new New Screen Payload
func NewNewScreenPayload(name string) Payload {
	return Payload{
		Type: "new_screen",
		Content: map[string]string {
			"name": name,
		},
	}
}
