package widget

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type Window struct {
	Title  string
	Height unit.Dp
	Width  unit.Dp

	Position f32.Point

	// Titlebar
	offset f32.Point
	start  f32.Point
	drag   gesture.Drag

	resizer *resizer
}

func NewWindow(title string, pos f32.Point) *Window {
	height := unit.Dp(500)
	width := unit.Dp(400)

	win := &Window{
		Title:    title,
		Height:   height,
		Width:    width,
		Position: pos,
		resizer:  &resizer{},
	}
	win.resizer = newResizer(win)
	return win
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	// TitleBar
	for _, e := range w.drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			w.start = e.Position
		case pointer.Drag:
			w.offset = e.Position.Sub(w.start)
		case pointer.Release:
			w.Position = w.Position.Add(w.offset)
			w.offset = f32.Point{}
		}
	}

	//w.resizer.update(gtx)

	m := op.Record(gtx.Ops)
	dims := widget(gtx)
	c := m.Stop()

	//w.resizer.layout(gtx, dims)

	defer op.Offset(w.Position.Round()).Push(gtx.Ops).Pop()

	// Resizer
	pad := gtx.Dp(unit.Dp(10))
	s := op.Offset(image.Pt(-pad, -pad)).Push(gtx.Ops)
	r := clip.Rect{
		Max: image.Pt(dims.Size.X+pad*2, dims.Size.Y+pad*2),
	}.Push(gtx.Ops)
	paint.Fill(gtx.Ops, color.NRGBA{R: 255, A: 100})
	r.Pop()
	s.Pop()

	defer op.Offset(w.offset.Round()).Push(gtx.Ops).Pop()
	c.Add(gtx.Ops)

	return dims
}

func (w *Window) TitleBar(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)

	if w.drag.Dragging() {
		defer op.Offset(w.offset.Round().Mul(-1)).Push(gtx.Ops).Pop()
	}
	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color.NRGBA{G: 255, A: 100})
	w.drag.Add(gtx.Ops)
	return dims
}
