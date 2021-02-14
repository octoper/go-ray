package payloads

func NewLogPayload(values ...interface{}) Payload {
	return Payload {
		Type: "log",
		Content: map[string]interface{} {
			"values": values,
		},
	}
}