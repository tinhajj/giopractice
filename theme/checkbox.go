package theme

import (
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type CheckboxStyle struct {
	b    *widget.Bool
	Size image.Point
}

func Checkbox(b *widget.Bool) *CheckboxStyle {
	return &CheckboxStyle{
		b:    b,
		Size: image.Point{X: 30, Y: 30},
	}
}

func (cs *CheckboxStyle) Layout(c layout.Context) layout.Dimensions {
	return widget.Border{Width: unit.Dp(3), Color: color.NRGBA{G: 100, A: 255}}.Layout(c, func(c layout.Context) layout.Dimensions {
		return cs.b.Layout(c, func(context layout.Context) layout.Dimensions {
			paint.FillShape(context.Ops, color.NRGBA{A: 100, G: 100}, clip.Rect{
				Min: image.Point{},
				Max: cs.Size,
			}.Op())

			if *cs.b.Value == true {
				paint.FillShape(context.Ops, color.NRGBA{
					R: 200,
					G: 0,
					B: 0,
					A: 100,
				}, clip.Rect{
					Min: image.Point{
						X: 0,
						Y: 0,
					},
					Max: cs.Size,
				}.Op())
			}

			return layout.Dimensions{Size: cs.Size}
		})
	})
}
