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

	ContentWidget layout.Widget

	offset f32.Point
	start  f32.Point

	click bool
	mask  any

	// Titlebar
	titleDrag gesture.Drag

	// Resizer
	resizeDrag   gesture.Drag
	resizeOffset image.Point
	startWidth   unit.Dp
	startHeight  unit.Dp
}

func NewWindow(title string, pos f32.Point, content layout.Widget) *Window {
	height := unit.Dp(500)
	width := unit.Dp(400)

	return &Window{
		Title:         title,
		Height:        height,
		Width:         width,
		Position:      pos,
		ContentWidget: content,
	}
}

func (w *Window) Clicked() bool {
	return w.click
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	// Active
	w.click = false
	for _, e := range gtx.Queue.Events(&w.click) {
		if x, ok := e.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
				w.click = true
			}
		}
	}

	// TitleBar
	for _, e := range w.titleDrag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
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

	// Resize
	for _, e := range w.resizeDrag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			w.start = e.Position
			w.startWidth = w.Width
			w.startHeight = w.Height
		case pointer.Drag:
			x := gtx.Metric.PxToDp(e.Position.Sub(w.start).Round().X)
			y := gtx.Metric.PxToDp(e.Position.Sub(w.start).Round().Y)

			// Clamp
			var width, height unit.Dp
			width = w.startWidth + x
			height = w.startHeight + y
			if width < unit.Dp(300) {
				width = unit.Dp(300)
			}
			if height < unit.Dp(300) {
				height = unit.Dp(300)
			}
			w.Width = width
			w.Height = height
		case pointer.Release:
			w.Position = w.Position.Add(w.offset)
			w.offset = f32.Point{}
		}
	}

	m := op.Record(gtx.Ops)
	dims := widget(gtx)
	c := m.Stop()

	defer op.Offset(w.Position.Round()).Push(gtx.Ops).Pop()

	// Resizer
	pad := gtx.Dp(unit.Dp(15))
	if !w.resizeDrag.Dragging() {
		w.resizeOffset = dims.Size
	}
	s := op.Offset(w.resizeOffset).Push(gtx.Ops)
	r := clip.Rect{
		Max: image.Pt(pad, pad),
	}.Push(gtx.Ops)
	w.resizeDrag.Add(gtx.Ops)
	pointer.CursorNorthWestSouthEastResize.Add(gtx.Ops)
	paint.Fill(gtx.Ops, color.NRGBA{R: 255, A: 100})
	r.Pop()
	s.Pop()

	defer op.Offset(w.offset.Round()).Push(gtx.Ops).Pop()

	// Mask
	rect := clip.Rect{Max: dims.Size}.Push(gtx.Ops)
	pointer.InputOp{
		Tag:   &w.mask,
		Types: AllPointers,
	}.Add(gtx.Ops)
	rect.Pop()

	// Draw the window
	c.Add(gtx.Ops)

	// Active Click
	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
	ps := pointer.PassOp{}.Push(gtx.Ops)
	pointer.InputOp{
		Tag:   &w.click,
		Types: pointer.Press | pointer.Release,
	}.Add(gtx.Ops)
	ps.Pop()

	return dims
}

func (w *Window) TitleBar(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)

	if w.titleDrag.Dragging() {
		defer op.Offset(w.offset.Round().Mul(-1)).Push(gtx.Ops).Pop()
	}
	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
	w.titleDrag.Add(gtx.Ops)
	return dims
}
