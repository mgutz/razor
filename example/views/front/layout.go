// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package front

import (
	"github.com/mgutz/razor"
)

// Layout is generated
func Layout(__data interface{}, __body *razor.SafeBuffer, __sections *razor.Sections) *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	data := __data.(razor.M)

	App := razor.App

	RenderBody := func() *razor.SafeBuffer {
		return __body
	}

	RenderSection := func(section string, required ...bool) *razor.SafeBuffer {
		sections := *__sections
		text := sections[section]
		isRequired := len(required) == 1 && required[0]
		if text == nil && isRequired {
			return razor.NewSafeBufferString("<div style='color:white; background-color: red'>SECTION " + section + " is required!<div>")
		}
		return text
	}
	__buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\" />\n    <title>")
	__buffer.WriteSafe(data["title"])
	__buffer.WriteString("</title>\n    <link rel=\"stylesheet\" href=\"/")
	__buffer.WriteSafe(App["version"])
	__buffer.WriteString("/css/style.css\">\n<body>\n    <div class=\"container\">\n      ")
	__buffer.WriteSafe(RenderBody())
	__buffer.WriteString("\n    </div>\n    ")
	__buffer.WriteSafe(data["footer"])
	__buffer.WriteString("\n    ")
	__buffer.WriteSafe(RenderSection("bodyFoot"))
	__buffer.WriteString("\n  </body>\n</html>")

	return __buffer
}
