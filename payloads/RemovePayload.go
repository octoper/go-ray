package payloads

// NewRemovePayload creates a new Remove Payload
func NewRemovePayload() Payload {
	return Payload {
		Type: "remove",
	}
}
