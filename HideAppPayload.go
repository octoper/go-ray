package ray

func NewHideAppPayload() Payload {
	return Payload{
		Type: "hide_app",
	}
}