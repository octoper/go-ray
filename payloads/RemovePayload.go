package payloads

func NewRemovePayload() Payload {
	return Payload{
		Type: "remove",
	}
}