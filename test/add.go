// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"cases/layout/base"
	"github.com/mgutz/razor"
)

// Add is generated
func Add(content string, err string) *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	locals := razor.Locals
	if locals != nil {
		// avoids not declared error if locals is not used
	}
	_buffer.WriteString("\n\n<link rel=\"stylesheet\" href=\"/css/bootstrap-datetimepicker.css\">\n\n<style>\n.row {\n	margin-top: 10px;\n}\n</style>\n\n<h2>日程登记</h2>\n\n<div class=\"container-fluid\">\n	<form method=\"POST\" action=\"\">\n	<div class=\"row\" >\n		<p class=\"bg-danger\">")
	_buffer.WriteSafe(err)
	_buffer.WriteString("</p>\n	</div>\n\n	<div class=\"row\">\n	内容:\n	<input type='text' class=\"form-control\" name=\"content\" value=\"")
	_buffer.WriteSafe(content)
	_buffer.WriteString("\"/>\n	</div>\n\n	<div class=\"row\">\n	开始时间:\n	<input type='text' class=\"datetimepicker form-control\" name=\"startTime\"/>\n	</div>\n\n	<div class=\"row\">\n	结束时间:\n	<input type='text' class=\"datetimepicker form-control\" name=\"endTime\"/>\n	</div>\n\n	<div class=\"row\">\n	日程指派:\n	<select name=\"appoint\">\n		<option>cheney</option>\n		<option>wuvist</option>\n	</select>\n	</div>\n\n	<div class=\"row\">\n	<input style=\"float:right\" type=\"submit\" value=\"保存\" class=\"btn btn-primary\"/>\n	</div>\n	</form>\n</div>")

	title := func() *razor.SafeBuffer {
		_buffer := razor.NewSafeBuffer()

		_buffer.WriteString("管理后台 - 添加日程")

		return _buffer
	}

	js := func() *razor.SafeBuffer {
		_buffer := razor.NewSafeBuffer()

		_buffer.WriteString("<script src=\"/js/moment.js\"></script>")

		_buffer.WriteString("<script src=\"/js/bootstrap-datetimepicker.js\"></script>")

		_buffer.WriteString("<script type=\"text/javascript\">\n	$(function () {\n		$(\".datetimepicker\").datetimepicker({\n			format: \"YYYY-MM-DD HH:mm\",\n			defaultDate: \"2014-05-01 00:00\",\n		})\n	});\n</script>")

		return _buffer
	}

	_sections := make(razor.Sections)
	_sections["title"] = title()
	_sections["js"] = js()

	return _buffer
}