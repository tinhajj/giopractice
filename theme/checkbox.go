package theme

import (
	"context"
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/font"
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
	return widget.Border{Width: unit.Dp(2), Color: color.NRGBA{G: 100, A: 255}}.Layout(c, func(c layout.Context) layout.Dimensions {
		return cs.b.Layout(c, func(context layout.Context) layout.Dimensions {
			paint.FillShape(context.Ops, color.NRGBA{A: 100, G: 255}, clip.Rect{
				Max: c.Constraints.Max,
			}.Op())

			if cs.b.Value == true {
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

type CheckBoxStyle struct {
	Checkbox        widget.Boolean
	Label           string
	TextSize        unit.Sp
	ForegroundColor color.NRGBA
	BackgroundColor color.NRGBA
	TextColor       color.NRGBA
}

func CheckBox(th *Theme, checkbox widget.Boolean, label string) CheckBoxStyle {
	return CheckBoxStyle{
		Checkbox:        checkbox,
		Label:           label,
		TextColor:       th.Palette.Foreground,
		ForegroundColor: th.Palette.Foreground,
		BackgroundColor: rgba(0),
		TextSize:        th.TextSize,
	}
}

func (c CheckBoxStyle) Layout(win *Window, gtx layout.Context) layout.Dimensions {
	defer rtrace.StartRegion(context.Background(), "theme.CheckBoxStyle.Layout").End()

	return c.Checkbox.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				sizeDp := gtx.Metric.SpToDp(c.TextSize)
				sizePx := gtx.Dp(sizeDp)

				ngtx := gtx
				ngtx.Constraints = layout.Exact(image.Pt(sizePx, sizePx))
				return widget.Border{
					Color: c.ForegroundColor,
					Width: 1,
				}.Layout(ngtx, func(gtx layout.Context) layout.Dimensions {
					paint.FillShape(gtx.Ops, c.BackgroundColor, clip.Rect{Max: gtx.Constraints.Min}.Op())
					if c.Checkbox.Get() {
						padding := gtx.Constraints.Min.X / 4
						if padding == 0 {
							padding = gtx.Dp(1)
						}
						minx := padding
						miny := minx
						maxx := gtx.Constraints.Min.X - padding
						maxy := maxx
						paint.FillShape(gtx.Ops, c.ForegroundColor, clip.Rect{Min: image.Pt(minx, miny), Max: image.Pt(maxx, maxy)}.Op())
					}

					return layout.Dimensions{Size: gtx.Constraints.Min}
				})
			}),

			layout.Rigid(layout.Spacer{Width: 3}.Layout),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widget.TextLine{Color: c.TextColor}.Layout(gtx, win.Theme.Shaper, font.Font{}, c.TextSize, c.Label)
			}),
		)
	})
}
