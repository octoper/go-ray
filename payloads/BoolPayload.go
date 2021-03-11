package payloads

// NewBoolPayload creates a new Bool Payload
func NewBoolPayload(bool bool) Payload {
	return NewCustomPayload(bool, "Boolean")
}
