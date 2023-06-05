package theme

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

func Button(clickable *widget.Clickable, label string) ButtonStyle {
	return ButtonStyle{
		Button: clickable,
		Label:  label,
		Font:   text.Font{},
	}
}

func (b ButtonStyle) Layout(gtx layout.Context) layout.Dimensions {
	return widget.Border{
		Color:        Theme.Black,
		CornerRadius: 0,
		Width:        unit.Dp(2),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			defer clip.UniformRRect(image.Rectangle{
				Min: image.Point{},
				Max: gtx.Constraints.Max,
			}, 2).Push(gtx.Ops).Pop()

			var c color.NRGBA

			if b.Button.Pressed() {
				c = color.NRGBA{R: 211, G: 211, B: 211, A: 255}
			} else {
				c = Theme.White
			}

			paint.Fill(gtx.Ops, c)

			colMacro := op.Record(gtx.Ops)
			paint.ColorOp{Color: Theme.Palette.Black}.Add(gtx.Ops)
			colOp := colMacro.Stop()

			return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return widget.Label{}.Layout(gtx, Theme.Shaper, b.Font, Theme.TextSize, b.Label, colOp)
			})
		})
	})
}

type ButtonStyle struct {
	Button *widget.Clickable
	Label  string
	Font   text.Font
}
