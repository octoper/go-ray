package payloads

func NewNotifyPayload(value string) Payload {
	return Payload{
		Type: "log",
		Content: map[string]interface{} {
			"value": value,
		},
	}
}