package ray

func NewClearAllPayload() Payload {
	return Payload{
		Type: "clear_all",
	}
}