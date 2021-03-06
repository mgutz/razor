// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
)

// Badtag is generated
func Badtag(w *gorazor.Widget) *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	if w.ErrorMsg != "" {

		__buffer.WriteString("<div class=\"form-group has-error\">\n	<div class=\"alert alert-danger\">")
		__buffer.WriteSafe(w.ErrorMsg)
		__buffer.WriteString("</div>")
	} else {

		__buffer.WriteString("<div class=\"form-group\">")
	}
	__buffer.WriteString("\n\n	<label for=\"")
	__buffer.WriteSafe(w.Name)
	__buffer.WriteString("\">")
	__buffer.WriteSafe(w.Label)
	__buffer.WriteString("</label>\n	<input type=\"text\" name=\"")
	__buffer.WriteSafe(w.Name)
	__buffer.WriteString("\" class=\"form-control\" id=\"")
	__buffer.WriteSafe(w.Name)
	__buffer.WriteString("\" placeholder=\"")
	__buffer.WriteSafe(w.PlaceHolder)
	__buffer.WriteString("\" value=\"")
	__buffer.WriteSafe(w.Value)
	__buffer.WriteString("\">\n</div>")

	return __buffer
}
