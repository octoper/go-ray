package payloads

type ColorPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewColorPayload(color string) *ColorPayload {
	return &ColorPayload{
		Type: "color",
		Content: map[string]string {
			"color": color,
		},
	}
}