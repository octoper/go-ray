package payloads

func NewHidePayload() Payload {
	return Payload{
		Type: "hide",
	}
}