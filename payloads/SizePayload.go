package payloads

func NewSizePayload(color string) Payload {
	return Payload{
		Type: "size",
		Content: map[string]string {
			"size": color,
		},
	}
}