package payloads

type ClearScreenPayload struct {
	Type string `json:"type"`
}

func NewClearScreenPayload() *ClearScreenPayload {
	return &ClearScreenPayload{
		Type: "new_screen",
	}
}