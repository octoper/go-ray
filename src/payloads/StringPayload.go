package payloads

func NewStringPayload(value string) Payload {
	return NewCustomPayload(value, "String")
}