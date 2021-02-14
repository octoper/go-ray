package payloads

func NewClearScreenPayload() Payload {
	return Payload{
		Type: "new_screen",
	}
}