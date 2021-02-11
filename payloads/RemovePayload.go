package payloads

type RemovePayload struct {
	Type string `json:"type"`
}

func NewRemovePayload() *RemovePayload {
	return &RemovePayload{
		Type: "remove",
	}
}