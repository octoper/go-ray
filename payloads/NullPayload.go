package payloads

func NewNullPayload() *CustomPayload {
	return NewCustomPayload(nil, "Null")
}