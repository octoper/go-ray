package payloads

// NewCreateLockPayload creates a new Create Lock Payload
func NewCreateLockPayload(name string) Payload {
	return Payload {
		Type: "create_lock",
		Content: map[string]string {
			"name": name,
		},
	}
}
