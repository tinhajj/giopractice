package widget

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type Window struct {
	Title    string
	Height   unit.Dp
	Width    unit.Dp
	Position f32.Point

	Resizer *Resizer
}

func NewWindow(title string) *Window {
	height := unit.Dp(500)
	width := unit.Dp(400)

	return &Window{
		Title:    title,
		Height:   height,
		Width:    width,
		Position: f32.Point{},
		Resizer:  NewResizer(),
	}
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	defer op.Offset(w.Position.Round()).Push(gtx.Ops).Pop()
	dims := w.Resizer.Layout(gtx, widget)

	return layout.Dimensions{
		Size:     dims.Size,
		Baseline: 0,
	}
}
