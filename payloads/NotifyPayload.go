package payloads

func NewNotifyPayload(value string) Payload {
	return Payload{
		Type: "notify",
		Content: map[string]interface{} {
			"value": value,
		},
	}
}