// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Section is generated
func Section() *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	_buffer.WriteString("\n\n<h1>Hello</h1>")

	side := func() *razor.SafeBuffer {
		_buffer := razor.NewSafeBuffer()

		_buffer.WriteString("<p>side</p>")

		return _buffer
	}

	_sections := make(razor.Sections)
	_sections["side"] = side()
	_buffer = SectionLayout(_buffer, &_sections)
	return _buffer
}
