// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor"
	"strconv"
	"zfw/models"
)

// Slashbug is generated
func Slashbug(objs ...*models.Widget) *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	locals := razor.Locals
	if locals != nil {
		// avoids not declared error if locals is not used
	}

	size := strconv.Itoa(12 / len(objs))

	return _buffer
}