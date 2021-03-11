package payloads

// Creates a new Bool Payload
func NewBoolPayload(bool bool) Payload {
	return NewCustomPayload(bool, "Boolean")
}
