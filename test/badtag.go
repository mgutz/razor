// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Badtag is generated
func Badtag(w *gorazor.Widget) *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	if w.ErrorMsg != "" {

		_buffer.WriteString("<div class=\"form-group has-error\">\n	<div class=\"alert alert-danger\">")
		_buffer.WriteSafe(w.ErrorMsg)
		_buffer.WriteString("</div>")
	} else {

		_buffer.WriteString("<div class=\"form-group\">")
	}
	_buffer.WriteString("\n\n	<label for=\"")
	_buffer.WriteSafe(w.Name)
	_buffer.WriteString("\">")
	_buffer.WriteSafe(w.Label)
	_buffer.WriteString("</label>\n	<input type=\"text\" name=\"")
	_buffer.WriteSafe(w.Name)
	_buffer.WriteString("\" class=\"form-control\" id=\"")
	_buffer.WriteSafe(w.Name)
	_buffer.WriteString("\" placeholder=\"")
	_buffer.WriteSafe(w.PlaceHolder)
	_buffer.WriteString("\" value=\"")
	_buffer.WriteSafe(w.Value)
	_buffer.WriteString("\">\n</div>")

	return _buffer
}
