package payloads

func NewColorPayload(color string) Payload {
	return Payload{
		Type: "color",
		Content: map[string]string {
			"color": color,
		},
	}
}