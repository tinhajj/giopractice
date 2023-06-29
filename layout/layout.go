package layout

import (
	"image"
	"ui/event"

	"gioui.org/layout"
	"gioui.org/op"
)

func Largest(gtx layout.Context, widgets ...layout.Widget) (int, layout.Dimensions) {
	largest := layout.Dimensions{}
	largestIndex := -1

	gtx.Queue = event.DummyQueue()
	m := op.Record(gtx.Ops)
	for i, w := range widgets {
		dim := w(gtx)

		if dim.Size.X > largest.Size.X {
			largest = dim
			largestIndex = i
		}
	}
	m.Stop()
	return largestIndex, largest
}

type ConstrainedLayout struct {
	Size image.Point
}

func (cl *ConstrainedLayout) Layout(c layout.Context, w layout.Widget) layout.Dimensions {
	c.Constraints.Max = cl.Size
	return w(c)
}
