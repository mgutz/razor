// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Email is generated
func Email() *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	__buffer.WriteString("<span>rememberingsteve@apple.com ")
	__buffer.WriteSafe(username)
	__buffer.WriteString("</span>")

	return __buffer
}
