// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Bug9 is generated
func Bug9() *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	locals := razor.Locals
	if locals != nil {
		// avoids not declared error if locals is not used
	}
	var l *Locale
	_buffer.WriteString("\n<span>")
	_buffer.WriteSafe(l.T(`for`))
	_buffer.WriteString("</span>")

	return _buffer
}