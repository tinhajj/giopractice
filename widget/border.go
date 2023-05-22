package widget

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

type OuterBorder struct {
	widget.Border
}

func (ob OuterBorder) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return ob.Border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(ob.Border.Width)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return w(gtx)
		})
	})
}
