package payloads

func NewNewScreenPayload(name string) Payload {
	return Payload{
		Type: "new_screen",
		Content: map[string]string {
			"name": name,
		},
	}
}