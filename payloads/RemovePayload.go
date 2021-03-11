package payloads

// Create a new Remove Payload
func NewRemovePayload() Payload {
	return Payload{
		Type: "remove",
	}
}
