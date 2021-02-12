package payloads

type NotifyPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewNotifyPayload(value string) *NotifyPayload {
	return &NotifyPayload{
		Type: "log",
		Content: map[string]interface{} {
			"value": value,
		},
	}
}