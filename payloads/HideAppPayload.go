package payloads

// Create a new Hide App Payload
func NewHideAppPayload() Payload {
	return Payload{
		Type: "hide_app",
	}
}
