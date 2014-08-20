package razor

import (
	"bytes"
	"html/template"
	"io"
)

var Empty = NewSafeBuffer()

type SafeBuffer struct {
	*bytes.Buffer
}

type Sections map[string]*SafeBuffer

func NewSafeBuffer() *SafeBuffer {
	return &SafeBuffer{Buffer: bytes.NewBuffer(nil)}
}

func NewSafeBufferString(s string) *SafeBuffer {
	return &SafeBuffer{Buffer: bytes.NewBufferString(s)}
}

func (self *SafeBuffer) WriteTo(w io.Writer) {
	self.Buffer.WriteTo(w)
}

func (self *SafeBuffer) WriteSafe(t interface{}) {
	switch v := t.(type) {
	case *SafeBuffer:
		if v != nil {
			self.Write(v.Bytes())
		}
	case bytes.Buffer:
		template.HTMLEscape(self.Buffer, v.Bytes())
	default:
		s := template.HTMLEscaper(v)
		if len(s) > 0 && s != "&lt;no value&gt;" {
			self.WriteString(s)
		}
	}
}
