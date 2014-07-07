// DO NOT EDIT! Auto-generated by github.com/mgutz/gorazor
package cases

import (
	"github.com/mgutz/razor/razor"
	. "kp/models"
)

func Msg() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	var u *User

	getName := func(u *User) string {
		return "(" + u.Name + ")"
	}

	var username string
	if u.Email != "" {
		username = getName(u) + "(" + u.Email + ")"
	}
	_buffer.WriteString("\n<div class=\"welcome\">\n<h4>Hello ")
	_buffer.WriteSafe(username)
	_buffer.WriteString("</h4>\n\n<div>")
	_buffer.WriteSafe((u.Intro))
	_buffer.WriteString("</div>\n</div>")

	return _buffer
}
