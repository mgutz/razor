// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor/razor"
)

// Layout is generated
func Layout() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	var body string
	var title string
	var side string
	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"utf-8\" />")
	_buffer.WriteSafe(title)
	_buffer.WriteString("\n</head>\n<body>\n<div>")
	_buffer.WriteSafe(body)
	_buffer.WriteString("</div>\n<div>")
	_buffer.WriteSafe(side)
	_buffer.WriteString("</div>\n</body>\n</html>")

	return _buffer
}
