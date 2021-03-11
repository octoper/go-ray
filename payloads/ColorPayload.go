package payloads

// Create a new Color Payload
func NewColorPayload(color string) Payload {
	return Payload{
		Type: "color",
		Content: map[string]string {
			"color": color,
		},
	}
}
