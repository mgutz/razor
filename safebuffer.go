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
	return &SafeBuffer{Buffer: new(bytes.Buffer)}
}

// NewSafeBufferString creates a new SafeBuffer from string.
func NewSafeBufferString(s string) *SafeBuffer {
	return &SafeBuffer{Buffer: bytes.NewBufferString(s)}
}

// WriteSafe writes the escaped value to the buffer.
func (sbuf *SafeBuffer) WriteSafe(t interface{}) {
	switch v := t.(type) {
	case *SafeBuffer:
		if v != nil {
			sbuf.Write(v.Bytes())
		}
	case bytes.Buffer:
		template.HTMLEscape(sbuf.Buffer, v.Bytes())
	default:
		s := template.HTMLEscaper(v)
		if len(s) > 0 && s != "&lt;no value&gt;" {
			sbuf.WriteString(s)
		}
	}
}

// RewriteString resets and writes s to the buffer.
func (sbuf *SafeBuffer) RewriteString(s string) {
	sbuf.Reset()
	sbuf.WriteString(s)
}

// Rewrite resets and writes bytes to the buffer.
func (sbuf *SafeBuffer) Rewrite(bytes []byte) {
	sbuf.Reset()
	sbuf.Write(bytes)
}
