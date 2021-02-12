package payloads

type CustomPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewCustomPayload(content interface{}, label string) *CustomPayload {
	return &CustomPayload{
		Type: "custom",
		Content: map[string]interface{} {
			"content": content,
			"label": label,
		},
	}
}