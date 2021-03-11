package payloads

// Create a new Custom Payload
func NewCustomPayload(content interface{}, label string) Payload {
	return Payload{
		Type: "custom",
		Content: map[string]interface{} {
			"content": content,
			"label": label,
		},
	}
}
