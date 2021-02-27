package payloads

func NewHtmlPayload(html string) Payload {
	return NewCustomPayload(html, "HTML")
}