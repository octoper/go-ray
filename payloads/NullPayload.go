package payloads

// Create a new Null Payload
func NewNullPayload() Payload {
	return NewCustomPayload(nil, "Null")
}
