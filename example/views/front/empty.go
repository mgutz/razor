// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package front

import (
	"github.com/mgutz/razor"
)

// Empty is generated
func Empty(__data interface{}) *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	__buffer.WriteString("\n\n<h1>Empty</h1>")

	__buffer = Layout("empty", __buffer, nil)
	return __buffer
}