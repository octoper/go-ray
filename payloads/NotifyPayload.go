package payloads

// Create a new Notify Payload
func NewNotifyPayload(value string) Payload {
	return Payload{
		Type: "notify",
		Content: map[string]interface{} {
			"value": value,
		},
	}
}
