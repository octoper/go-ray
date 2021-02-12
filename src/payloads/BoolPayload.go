package payloads

func NewBoolPayload(bool bool) *CustomPayload {
	return NewCustomPayload(bool, "Boolean")
}