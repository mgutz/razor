// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package views

import (
	"github.com/mgutz/razor/example/models"
	"github.com/mgutz/razor/example/views/layout"
	"github.com/mgutz/razor/html"
	"github.com/mgutz/razor/razor"
)

//Index is generated
func Index(user *models.User) razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	title := "Cool Site"
	_buffer.WriteSafe(html.Raw(`
<div style='background-color: #00F; color: #fff;'>
Why so blue?
</div>
`))
	_buffer.WriteString("\n\n<p>Escaped: ")
	_buffer.WriteSafe(UnsafeHello(user.Name))
	_buffer.WriteString("</p>\n<p>Unescaped: ")
	_buffer.WriteSafe(SafeHello(user.Name))
	_buffer.WriteString("</p>")
	_buffer.WriteSafe(html.Raw("<h2>Heading 2</h2>"))

	js := func() razor.SafeBuffer {
		_buffer := razor.NewSafeBuffer()

		_buffer.WriteString("<script>\n    alert('Hello, ")
		_buffer.WriteSafe(user.Name)
		_buffer.WriteString("');\n  </script>")

		return _buffer
	}

	_buffer = layout.Default(_buffer, js(), title)
	return _buffer
}
