// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package benchfiles

import (
	"github.com/mgutz/razor"
)

// Layout is generated
func Layout(u *User, body *razor.SafeBuffer, psections *razor.Sections) *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()

	RenderBody := func() *razor.SafeBuffer {
		return body
	}
	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n    <body>\n\n      <header>\n        <h1>")
	_buffer.WriteSafe(u.FirstName)
	_buffer.WriteString(" ")
	_buffer.WriteSafe(u.LastName)
	_buffer.WriteString(" ")
	_buffer.WriteSafe(u.Age)
	_buffer.WriteString("</h1>\n        <nav>\n          <ul>\n            ")
	for _, hobby := range u.Hobbies {

		_buffer.WriteString("<li><a href=\"hobby\">")
		_buffer.WriteSafe(hobby)
		_buffer.WriteString("</a></li>")

	}
	_buffer.WriteString("\n          </ul>\n        </nav>\n      </header>\n\n      <article id=\"content\">\n        ")
	_buffer.WriteSafe(RenderBody())
	_buffer.WriteString("\n      </article>\n\n      <footer>\n        &copy; Copyright 2013 by golang-samples.\n      </footer>\n    </body>\n</html>")

	return _buffer
}