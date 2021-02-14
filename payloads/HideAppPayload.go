package payloads

func NewHideAppPayload() Payload {
	return Payload{
		Type: "hide_app",
	}
}