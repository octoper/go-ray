package payloads

type ShowAppPayload struct {
	Type string `json:"type"`
}

func NewShowAppPayload() *ShowAppPayload {
	return &ShowAppPayload{
		Type: "show_app",
	}
}