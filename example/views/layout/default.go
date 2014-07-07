// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package layout

import (
	"github.com/mgutz/razor/razor"
)

//Default is generated
func Default(body, js razor.SafeBuffer, title string) razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\" />\n    <title>")
	_buffer.WriteSafe(title)
	_buffer.WriteString("</title>\n    <link rel=\"stylesheet\" href=\"/{{version}}/css/style.css\">\n</head>\n<body>\n    <div class=\"container\">")
	_buffer.WriteSafe(body)
	_buffer.WriteString("</div>\n    ")
	_buffer.WriteSafe(js)
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer
}
