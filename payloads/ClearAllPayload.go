package payloads

// Create a new Cleal All Payload
func NewClearAllPayload() Payload {
	return Payload{
		Type: "clear_all",
	}
}
