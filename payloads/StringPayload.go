package payloads

// NewStringPayload creates a new String Payload
func NewStringPayload(value string) Payload {
	return NewCustomPayload(value, "String")
}
