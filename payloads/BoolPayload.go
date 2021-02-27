package payloads

func NewBoolPayload(bool bool) Payload {
	return NewCustomPayload(bool, "Boolean")
}