// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package benchfiles

import (
	"github.com/mgutz/razor"
)

// Index is generated
func Index(__data interface{}) *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	u := __data.(*User)
	__buffer.WriteString("\n<h1>Hoge</h1>\n<p>\n  I'm Hoge.\n</p>")

	__scripts := func() *razor.SafeBuffer {
		__buffer := razor.NewSafeBuffer()

		__buffer.WriteString("<script>\n    console.log('foo');\n  </script>")

		return __buffer
	}

	__sections := make(razor.Sections)
	__sections["scripts"] = __scripts()
	__buffer = Layout(u, __buffer, &__sections)
	return __buffer
}
