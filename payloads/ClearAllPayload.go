package payloads

type ClearAllPayload struct {
	Type string `json:"type"`
	Content string `json:"content"`
}

func NewClearAllPayload() *ClearAllPayload {
	return &ClearAllPayload{
		Type: "clear_all",
	}
}