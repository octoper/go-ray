package payloads

func NewShowAppPayload() Payload {
	return Payload{
		Type: "show_app",
	}
}