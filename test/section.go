// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Section is generated
func Section() *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	__buffer.WriteString("\n\n<h1>Hello</h1>")

	__side := func() *razor.SafeBuffer {
		__buffer := razor.NewSafeBuffer()

		__buffer.WriteString("<p>side</p>")

		return __buffer
	}

	__sections := make(razor.Sections)
	__sections["side"] = __side()
	__buffer = SectionLayout(__buffer, &__sections)
	return __buffer
}
