package payloads

// NewNullPayload creates a new Null Payload
func NewNullPayload() Payload {
	return NewCustomPayload(nil, "Null")
}
