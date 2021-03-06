package payloads

func NewCreateLockPayload(name string) Payload {
	return Payload{
		Type: "create_lock",
		Content: map[string]string {
			"name": name,
		},
	}
}