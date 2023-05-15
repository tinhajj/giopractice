package theme

import (
	"image/color"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
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
	return widget.Border{Width: unit.Dp(3), Color: color.NRGBA{G: 100, A: 255}}.Layout(c, func(c layout.Context) layout.Dimensions {
		return cs.b.Layout(c, func(context layout.Context) layout.Dimensions {
			paint.FillShape(context.Ops, color.NRGBA{A: 100, G: 100}, clip.Rect{
				Max: c.Constraints.Max,
			}.Op())

			if *cs.b.Value == true {
				paint.FillShape(context.Ops, color.NRGBA{
					R: 200,
					A: 100,
				}, clip.Rect{
					Max: c.Constraints.Max,
				}.Op())
			}

			return layout.Dimensions{Size: c.Constraints.Max}
		})
	})
}
