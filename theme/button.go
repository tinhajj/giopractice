package theme

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Button(th *material.Theme, clickable *widget.Clickable, label string) ButtonStyle {
	return ButtonStyle{
		Button: clickable,
		Label:  label,
		Theme:  th,
	}
}

func (b ButtonStyle) Layout(gtx layout.Context) layout.Dimensions {
	return b.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		defer clip.UniformRRect(image.Rectangle{
			Min: image.Point{},
			Max: gtx.Constraints.Max,
		}, 2).Push(gtx.Ops).Pop()

		var c color.NRGBA

		if b.Button.Pressed() {
			c = color.NRGBA{R: 0, G: 0, B: 255, A: 150}
		} else {
			c = color.NRGBA{R: 0, G: 0, B: 255, A: 255}
		}

		paint.Fill(gtx.Ops, c)

		colMacro := op.Record(gtx.Ops)
		paint.ColorOp{Color: color.NRGBA{R: 255, A: 255}}.Add(gtx.Ops)
		colOp := colMacro.Stop()

		return layout.Inset{Top: 20, Bottom: 20, Left: 20, Right: 20}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return widget.Label{}.Layout(gtx, b.Theme.Shaper, b.Font, 30, b.Label, colOp)
		})
	})
}

type ButtonStyle struct {
	Button *widget.Clickable
	Label  string
	Theme  *material.Theme
	Font   text.Font
}
