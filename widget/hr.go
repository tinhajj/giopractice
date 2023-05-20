package widget

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type HR struct {
	Width unit.Dp
}

func (h HR) Layout(gtx layout.Context) layout.Dimensions {
	var path clip.Path

	path.Begin(gtx.Ops)
	path.MoveTo(f32.Pt(0, 0))
	path.LineTo(f32.Pt(float32(gtx.Constraints.Max.X), 0))
	path.Close()

	paint.FillShape(gtx.Ops, color.NRGBA{G: 255, A: 255},
		clip.Stroke{
			Path:  path.End(),
			Width: float32(gtx.Dp(h.Width)),
		}.Op())
	return layout.Dimensions{
		Size: image.Point{X: gtx.Constraints.Max.X, Y: gtx.Dp(h.Width)},
	}
}
