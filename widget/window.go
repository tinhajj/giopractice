package widget

import (
	"image"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type Window struct {
	Title  string
	Height unit.Dp
	Width  unit.Dp

	position f32.Point

	Resizer *Resizer
}

func NewWindow(title string) *Window {
	height := unit.Dp(500)
	width := unit.Dp(400)

	return &Window{
		Title:    title,
		Height:   height,
		Width:    width,
		position: f32.Point{},
		Resizer:  NewResizer(),
	}
}

func (w *Window) Position(p f32.Point) {
	w.position = p
	w.Resizer.position = p
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	w.Resizer.Layout(gtx, layout.Dimensions{
		Size: image.Point{X: gtx.Dp(w.Width), Y: gtx.Dp(w.Height)},
	})

	defer op.Offset(w.position.Round()).Push(gtx.Ops).Pop()

	if ok, rb := w.Resizer.Dragging(); ok {
		if rb.Direction == Vertical {
			w.Height = unit.Dp(rb.Offset.Round().X)
		}
		w.position = w.position.Add(rb.Offset)
	}

	return widget(gtx)
}
