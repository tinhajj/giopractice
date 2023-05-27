package widget

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type Window struct {
	Title  string
	Height unit.Dp
	Width  unit.Dp

	Position f32.Point
	offset   f32.Point

	Resizer *Resizer
}

func NewWindow(title string) *Window {
	height := unit.Dp(500)
	width := unit.Dp(400)

	win := &Window{
		Title:    title,
		Height:   height,
		Width:    width,
		Position: f32.Point{},
	}
	win.Resizer = NewResizer(win)
	return win
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	w.Resizer.Layout(gtx)

	defer op.Offset(w.Position.Round()).Push(gtx.Ops).Pop()

	return widget(gtx)
}
