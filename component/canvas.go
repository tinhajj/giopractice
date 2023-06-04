package component

import (
	"image"
	"ui/theme"
	"ui/widget"

	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Canvas struct {
	Windows []*widget.Window

	GridSize int

	origin f32.Point

	drag   gesture.Drag
	offset f32.Point
	start  f32.Point
}

func (c *Canvas) Layout(gtx layout.Context) layout.Dimensions {
	gridSize := c.GridSize
	if gridSize == 0 {
		gridSize = 48
	}

	for _, e := range c.drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			c.start = e.Position
		case pointer.Drag:
			c.offset = e.Position.Sub(c.start)
		case pointer.Release:
			c.origin = c.origin.Add(c.offset)
			c.offset = f32.Point{}
		}
	}

	paint.Fill(gtx.Ops, theme.Theme.Bg)
	s := clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)
	c.drag.Add(gtx.Ops)
	s.Pop()

	var path clip.Path

	col := theme.Theme.LightGray
	col.A = 30

	// Vertical Line
	m := op.Record(gtx.Ops)
	path.Begin(gtx.Ops)
	path.Line(f32.Pt(0, float32(gtx.Constraints.Max.Y)))
	paint.FillShape(gtx.Ops, col, clip.Stroke{
		Path:  path.End(),
		Width: 1,
	}.Op())
	vert := m.Stop()

	// Horizontal Line
	m = op.Record(gtx.Ops)
	path.Begin(gtx.Ops)
	path.Line(f32.Pt(float32(gtx.Constraints.Max.X), 0))
	paint.FillShape(gtx.Ops, col, clip.Stroke{
		Path:  path.End(),
		Width: 1,
	}.Op())
	horiz := m.Stop()

	cord := c.origin.Round().Add(c.offset.Round())
	shift := image.Pt(0, 0)

	// Vertical Line
	initial := f32.Pt(float32(cord.X%gridSize), 0).Round()
	shift = shift.Add(initial)

	op.Offset(initial).Push(gtx.Ops)
	vert.Add(gtx.Ops)
	for shift.X < gtx.Constraints.Max.X {
		op.Offset(image.Pt(gridSize, 0)).Push(gtx.Ops)
		vert.Add(gtx.Ops)

		shift = shift.Add(image.Pt(gridSize, 0))
	}
	op.Offset(shift.Mul(-1)).Add(gtx.Ops)

	// Horizontal Line
	shift = image.Pt(0, 0)
	initial = f32.Pt(0, float32(cord.Y%gridSize)).Round()
	shift = shift.Add(initial)

	op.Offset(initial).Push(gtx.Ops)
	horiz.Add(gtx.Ops)
	for shift.Y < gtx.Constraints.Max.Y {
		op.Offset(image.Pt(0, gridSize)).Push(gtx.Ops)
		horiz.Add(gtx.Ops)

		shift = shift.Add(image.Pt(0, gridSize))
	}
	op.Offset(shift.Mul(-1)).Add(gtx.Ops)

	defer op.Offset(c.offset.Round()).Push(gtx.Ops).Pop()
	defer op.Offset(c.origin.Round()).Push(gtx.Ops).Pop()

	for i, w := range c.Windows {
		if w.Clicked() {
			op.InvalidateOp{}.Add(gtx.Ops)
			c.BringIndexedWindowToFront(i)
		}
		theme.Window(w).Layout(gtx)
	}

	return layout.Dimensions{
		Size: gtx.Constraints.Max,
	}
}

// BringIndexedWindowToFront paints the given window last so it will show on top of all other windows
func (c *Canvas) BringIndexedWindowToFront(i int) {
	if i < 0 || i >= len(c.Windows) {
		return
	}

	front := c.Windows[i]
	windows := append(c.Windows[:i], c.Windows[i+1:]...)
	c.Windows = append(windows, front)

	return
}
