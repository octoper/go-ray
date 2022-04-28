package payloads

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gomarkdown/markdown"
)

// NewDumpPayload creates a new Dump Payload
func NewDumpPayload(value interface{}) Payload {
	style := `
	<style>
		.go-dump pre {
			position:relative;
			overflow-x: auto;
			width: 100%;
			padding: 10px 10px;
			height: auto;
		}
	</style>`
	md := []byte("``` \n" + spew.Sdump(value) + "\n```")
	output := markdown.ToHTML(md, nil, nil)
	return NewCustomPayload(style+`<div class="go-dump">`+string(output)+"</div>", "")
}
