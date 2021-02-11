package payloads

type NewScreenPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewNewScreenPayload(name string) *NewScreenPayload {
	return &NewScreenPayload{
		Type: "new_screen",
		Content: map[string]string {
			"name": name,
		},
	}
}