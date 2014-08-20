// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package layout

import (
	"github.com/mgutz/razor/example/shared"
	"github.com/mgutz/razor/razor"
)

// Default is generated
func Default(body *razor.SafeBuffer, sections razor.Sections, data razor.ViewData) *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	locals := razor.Locals().(*shared.Locals)
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\" />\n    <title>")
	_buffer.WriteSafe(data["title"])
	_buffer.WriteString("</title>\n    <link rel=\"stylesheet\" href=\"/")
	_buffer.WriteSafe(locals.Version)
	_buffer.WriteString("/css/style.css\">\n    ")
	_buffer.WriteSafe(sections["headFoot"])
	_buffer.WriteString("\n</head>\n<body>\n    <div class=\"container\">")
	_buffer.WriteSafe(body)
	_buffer.WriteString("</div>\n    ")
	_buffer.WriteSafe(data["footer"])
	_buffer.WriteString("\n    ")
	_buffer.WriteSafe(sections["bodyFoot"])
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer
}
