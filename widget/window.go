package widget

import (
	"fmt"
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

	resizer  *resizer
	TitleBar *TitleBar
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
	win.TitleBar = &TitleBar{
		Window: win,
	}
	return win
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	w.resizer.update(gtx)
	w.TitleBar.update(gtx)

	m := op.Record(gtx.Ops)
	dims := widget(gtx)
	c := m.Stop()

	w.resizer.layout(gtx, dims)

	defer op.Offset(w.Position.Round()).Push(gtx.Ops).Pop()
	c.Add(gtx.Ops)

	return dims
}

type TitleBar struct {
	Window *Window

	startWindow       Window
	startDragPosition f32.Point

	offset f32.Point

	drag gesture.Drag
}

func (t *TitleBar) update(gtx layout.Context) {
	for _, e := range t.drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			t.startDragPosition = e.Position
			t.startWindow = *t.Window
		case pointer.Drag:
			t.offset = e.Position.Sub(t.startDragPosition)
			t.Window.Position = t.startWindow.Position.Add(t.offset)
		case pointer.Release:
			t.offset = f32.Point{}
		}
	}
}

func (t *TitleBar) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)

	if t.drag.Dragging() {
		fmt.Println(t.offset.Round().Mul(-1))
		defer op.Offset(t.offset.Round().Mul(-1)).Push(gtx.Ops).Pop()
	}
	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color.NRGBA{G: 255, A: 100})
	t.drag.Add(gtx.Ops)
	return dims
}
