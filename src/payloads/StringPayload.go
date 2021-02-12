package payloads

func NewStringPayload(value string) *CustomPayload {
	return NewCustomPayload(value, "String")
}