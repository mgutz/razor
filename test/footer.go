// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Footer is generated
func Footer() *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	locals := razor.Locals
	if locals != nil {
		// avoids not declared error if locals is not used
	}
	_buffer.WriteString("<div>copyright 2014</div>")

	return _buffer
}