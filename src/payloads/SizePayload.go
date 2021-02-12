package payloads

type SizePayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewSizePayload(color string) *SizePayload {
	return &SizePayload{
		Type: "size",
		Content: map[string]string {
			"size": color,
		},
	}
}