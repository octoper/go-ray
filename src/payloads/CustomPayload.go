package payloads

func NewCustomPayload(content interface{}, label string) Payload {
	return Payload{
		Type: "custom",
		Content: map[string]interface{} {
			"content": content,
			"label": label,
		},
	}
}