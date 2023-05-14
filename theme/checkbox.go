package theme

import (
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type CheckboxStyle struct {
	b *widget.Bool
}

func Checkbox(b *widget.Bool) *CheckboxStyle {
	return &CheckboxStyle{
		b: b,
	}
}

func (cs *CheckboxStyle) Layout(c layout.Context) layout.Dimensions {
	return cs.b.Layout(c, func(context layout.Context) layout.Dimensions {
		paint.FillShape(context.Ops, color.NRGBA{A: 100, G: 100}, clip.Rect{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{30, 30},
		}.Op())

		if cs.b.Value == true {
			paint.FillShape(context.Ops, color.NRGBA{
				R: 200,
				G: 0,
				B: 0,
				A: 100,
			}, clip.Rect{
				Min: image.Point{},
				Max: image.Point{30, 30},
			}.Op())
		}

		return layout.Dimensions{Size: image.Point{30, 30}}
	})
}
