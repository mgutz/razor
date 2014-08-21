package html

import (
	"bytes"
	"fmt"

	"github.com/mgutz/razor"
)

// Raw insert unescaped HTML markup.
//
//      @html.Raw("<div>foo</div>")
func Raw(markup interface{}) *razor.SafeBuffer {
	buffer := razor.NewSafeBuffer()
	switch v := markup.(type) {
	case *razor.SafeBuffer:
		if v != nil {
			buffer.Write(v.Bytes())
		}
	case bytes.Buffer:
		buffer.Write(v.Bytes())
	default:
		buffer.WriteString(fmt.Sprint(markup))
	}
	return buffer
}
