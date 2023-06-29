package theme

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

func Label(gtx layout.Context, size unit.Sp, txt string) layout.Dimensions {
	textColorMacro := op.Record(gtx.Ops)
	paint.ColorOp{Color: Theme.Black}.Add(gtx.Ops)
	textColor := textColorMacro.Stop()

	gtx.Constraints.Min.X = gtx.Constraints.Max.X

	label := widget.Label{Alignment: text.Start}

	return label.Layout(gtx, Theme.Shaper, Theme.Font, size, txt, textColor)
}
