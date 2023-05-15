package layout

import (
	"image"

	"gioui.org/layout"
)

type ConstrainedLayout struct {
	Size image.Point
}

func (cl *ConstrainedLayout) Layout(c layout.Context, w layout.Widget) layout.Dimensions {
	c.Constraints.Max = cl.Size
	return w(c)
}
