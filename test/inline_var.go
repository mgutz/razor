// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
	"github.com/sunfmin/gorazortests/models"
)

// Inline_var is generated
func Inline_var() *razor.SafeBuffer {
	__buffer := razor.NewSafeBuffer()
	__buffer.WriteString("\n\n<body>")
	__buffer.WriteSafe(Hello("Felix Sun", "h1", 30, &models.Author{"Van", 20}, 10))
	__buffer.WriteString("\n</body>")

	return __buffer
}
