// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
	"kp/models"
	"tpl/admin/layout/base"
)

// Edit is generated
func Edit() *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	var u *models.User
	_buffer.WriteString("\n<div style=\"width: 500px\">\n<form role=\"form\">\n  <div class=\"form-group\">\n    <label for=\"exampleInputEmail1\">名字</label>\n    <input type=\"email\" class=\"form-control\" id=\"exampleInputEmail1\" placeholder=\"Enter email\" value=\"")
	_buffer.WriteSafe(u.Name)
	_buffer.WriteString("\">\n  </div>\n  <div class=\"form-group\">\n    <label for=\"exampleInputPassword1\">电邮</label>\n    <input type=\"email\" class=\"form-control\" id=\"exampleInputPassword1\" placeholder=\"电邮\" value=\"")
	_buffer.WriteSafe(u.Email)
	_buffer.WriteString("\">\n  </div>\n  <button type=\"submit\" class=\"btn btn-primary\">保存</button>\n  <a href=\"/admin/user\" class=\"btn btn-default pull-right\">返回</a>\n</form>\n</div>")

	title := func() *razor.SafeBuffer {
		_buffer := razor.NewSafeBuffer()

		_buffer.WriteString("用户管理")

		return _buffer
	}

	return _buffer
}
