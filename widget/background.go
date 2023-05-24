package widget

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Background struct {
	Color color.NRGBA
}

func (b Background) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, b.Color)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return w(gtx)
		}),
	)
}
