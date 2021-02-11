package payloads

type HideAppPayload struct {
	Type string `json:"type"`
}

func NewHideAppPayload() *HideAppPayload {
	return &HideAppPayload{
		Type: "hide_app",
	}
}