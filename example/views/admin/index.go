// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package admin

import (
	"github.com/mgutz/razor"
	"github.com/mgutz/razor/example/models"
	"github.com/mgutz/razor/example/views"
)

// Index is generated
func Index(__data interface{}) *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	user := __data.(*models.User)
	data := razor.M{
		"title": "Razor + Go = love",
	}

	__body := func() *razor.SafeBuffer {
		__buffer := razor.NewSafeBuffer()
		views.Heading2("Admin area")

		__buffer.WriteString("Use foo sections for body to be consisten and it simplifies code")

		return __buffer
	}

	__bodyFoot := func() *razor.SafeBuffer {
		__buffer := razor.NewSafeBuffer()

		__buffer.WriteString("<!-- <script> -->\n  <!--   alert('Hello, ")
		__buffer.WriteSafe(user.Name)
		__buffer.WriteString("'); -->\n  <!-- </script>")

		__buffer.WriteString("-->")
		return __buffer
	}

	__sections := make(razor.Sections)
	__sections["body"] = __body()
	__sections["bodyFoot"] = __bodyFoot()
	__buffer = Layout(data, __buffer, &__sections)
	return __buffer
}
