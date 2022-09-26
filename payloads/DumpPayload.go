package payloads

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gomarkdown/markdown"
)

// NewDumpPayload creates a new Dump Payload
func NewDumpPayload(value interface{}) Payload {
	md := []byte("``` \n" + spew.Sdump(value) + "\n```")
	output := markdown.ToHTML(md, nil, nil)
	styledOutput := strings.Replace(string(output), "<pre><code>", strings.Join([]string{
		`<pre class="relative overflow-x-auto w-full p-5 h-auto bg-gray-100 dark:bg-gray-800">`,
		`<code class="h-auto">`,
	}, ""), 1)
	return NewCustomPayload(styledOutput, "")
}
