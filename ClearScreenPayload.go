package ray

func NewClearScreenPayload() Payload {
	return Payload{
		Type: "new_screen",
		Content: "",
	}
}