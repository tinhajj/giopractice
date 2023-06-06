package theme

import (
	"image"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type CheckboxStyle struct {
	Label string

	b *widget.Bool
}

func Checkbox(b *widget.Bool, label string) *CheckboxStyle {
	return &CheckboxStyle{
		Label: label,
		b:     b,
	}
}

func (c CheckboxStyle) Layout(gtx layout.Context) layout.Dimensions {
	return c.b.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				sizeDp := gtx.Metric.SpToDp(Theme.TextSize)
				sizePx := gtx.Dp(sizeDp)

				ngtx := gtx
				ngtx.Constraints = layout.Exact(image.Pt(sizePx, sizePx))
				return widget.Border{
					Color: Theme.Black,
					Width: 1,
				}.Layout(ngtx, func(gtx layout.Context) layout.Dimensions {
					paint.FillShape(gtx.Ops, Theme.Yellow, clip.Rect{Max: gtx.Constraints.Min}.Op())
					if c.b.Value {
						padding := gtx.Constraints.Min.X / 5
						if padding == 0 {
							padding = gtx.Dp(1)
						}
						minx := padding
						miny := minx
						maxx := gtx.Constraints.Min.X - padding
						maxy := maxx
						paint.FillShape(gtx.Ops, Theme.Black, clip.Rect{Min: image.Pt(minx, miny), Max: image.Pt(maxx, maxy)}.Op())
					}

					return layout.Dimensions{Size: gtx.Constraints.Min}
				})
			}),

			layout.Rigid(layout.Spacer{Width: 3}.Layout),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				colMacro := op.Record(gtx.Ops)
				paint.ColorOp{Color: Theme.Palette.Black}.Add(gtx.Ops)
				colOp := colMacro.Stop()
				return widget.Label{}.Layout(gtx, Theme.Shaper, Theme.Font, Theme.TextSize, c.Label, colOp)
			}),
		)
	})
}
