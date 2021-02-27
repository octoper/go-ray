package payloads

func NewNullPayload() Payload {
	return NewCustomPayload(nil, "Null")
}