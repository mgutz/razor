// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"dm"
	"github.com/mgutz/razor/razor"
	"zfw/models"
	. "zfw/tplhelper"
)

// Scopebug is generated
func Scopebug() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	var obj *models.Widget
	{
		if 1 == 2 {
		} else {
			values := []int{}
			for _, v := range values {
				if v, ok := v.(type); ok {

					_buffer.WriteString("<a>\n					")
					for _, v := range values {
					}
					_buffer.WriteString("\n				</a>")

				} else {

				}
			}
		}
	}

	return _buffer
}
