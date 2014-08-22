package razor

import (
	"bytes"
	"html/template"
	//"io"
)

// SafeBuffer is a bufffer which mitigates HTML escaping and raw values.
type SafeBuffer struct {
	*bytes.Buffer
}

// NewSafeBuffer creates a new SafeBuffer.
func NewSafeBuffer() *SafeBuffer {
	return &SafeBuffer{Buffer: bytes.NewBuffer(nil)}
}

// NewSafeBufferString creates a new SafeBuffer from string.
func NewSafeBufferString(s string) *SafeBuffer {
	return &SafeBuffer{Buffer: bytes.NewBufferString(s)}
}

// WriteSafe writes the escaped value to the buffer.
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
