package payloads

func NewHtmlPayload(html string) *CustomPayload {
	return NewCustomPayload(html, "HTML")
}