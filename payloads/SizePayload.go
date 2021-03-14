package payloads

// NewSizePayload creates a new Size Payload
func NewSizePayload(color string) Payload{
	return Payload{
		Type: "size",
		Content: map[string]string {
			"size": color,
		},
	}
}
