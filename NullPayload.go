package ray

func NewNullPayload() Payload {
	return NewCustomPayload(nil, "Null")
}