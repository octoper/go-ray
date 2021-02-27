package ray

func NewRemovePayload() Payload {
	return Payload{
		Type: "remove",
	}
}