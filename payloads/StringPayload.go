package payloads

// Create a new String Payload
func NewStringPayload(value string) Payload {
	return NewCustomPayload(value, "String")
}
