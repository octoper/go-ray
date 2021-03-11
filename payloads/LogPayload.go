package payloads

// NewLogPayload creates a new Log Payload
func NewLogPayload(values ...interface{}) Payload {
	return Payload {
		Type: "log",
		Content: map[string]interface{} {
			"values": values,
		},
	}
}
