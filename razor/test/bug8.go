// DO NOT EDIT! Auto-generated by github.com/mgutz/gorazor
package cases

import (
	"github.com/mgutz/razor/razor"
)

func Bug8() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	var l *Locale
	_buffer.WriteString("\n<span>")
	_buffer.WriteSafe(l.T("for"))
	_buffer.WriteString("</span>")

	return _buffer
}
