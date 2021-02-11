package payloads

type HidePayload struct {
	Type string `json:"type"`
}

func NewHidePayload() *HidePayload {
	return &HidePayload{
		Type: "hide",
	}
}