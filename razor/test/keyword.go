// DO NOT EDIT! Auto-generated by github.com/mgutz/gorazor
package cases

import (
	"github.com/mgutz/razor/razor"
)

func Keyword() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	_buffer.WriteString("BLK(<span>rememberingsteve@apple.com ")
	_buffer.WriteSafe(username)
	_buffer.WriteString("</span>)BLK")

	return _buffer
}
