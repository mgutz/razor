package razor

import (
	"bytes"
	//"log"
	"testing"
)

func TestSafeBuffer(t *testing.T) {
	buffer := NewSafeBuffer()

	buffer.WriteSafe("<")
	if buffer.String() != "&lt;" {
		t.Error("HTML entities should be escaped")
	}
	buffer.Reset()

	var bb bytes.Buffer
	bb.WriteString(">")
	buffer.WriteSafe(bb)
	if buffer.String() != "&gt;" {
		t.Error("HTML entities should be escaped from bytes.Buffer")
	}
	buffer.Reset()

	other := NewSafeBuffer()
	other.WriteString("<html>")
	buffer.WriteSafe(other)
	if buffer.String() != "<html>" {
		t.Error("WriteSafe should not escape SafeBuffer")
	}
	buffer.Reset()

	bb.Reset()
	bb.WriteString("<html>")
	buffer.WriteSafe(bb)
	if buffer.String() != "&lt;html&gt;" {
		t.Error("WriteSafe should escape bytes.Buffer")
	}
	buffer.Reset()
}
