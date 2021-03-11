package payloads

// Create a new HTML Payload
func NewHtmlPayload(html string) Payload {
	return NewCustomPayload(html, "HTML")
}
