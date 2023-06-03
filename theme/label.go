package theme

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

func Label(gtx layout.Context, size unit.Sp, txt string) layout.Dimensions {
	textColorMacro := op.Record(gtx.Ops)
	paint.ColorOp{Color: color.NRGBA{R: 0, B: 0, G: 0, A: 255}}.Add(gtx.Ops)
	textColor := textColorMacro.Stop()

	gtx.Constraints.Min.X = gtx.Constraints.Max.X

	label := widget.Label{Alignment: text.Start}

	return label.Layout(gtx, Theme.Shaper, Theme.Font, size, txt, textColor)
}
