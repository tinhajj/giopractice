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

type Resizer struct {
	Top    *ResizeBar
	Bottom *ResizeBar
	Left   *ResizeBar
	Right  *ResizeBar

	Thickness unit.Dp

	position f32.Point

	start  f32.Point
	offset f32.Point
}

func NewResizer() *Resizer {
	return &Resizer{
		Top:       &ResizeBar{Direction: Vertical, Positioner: true},
		Bottom:    &ResizeBar{Direction: Vertical, Positioner: false},
		Left:      &ResizeBar{Direction: Horizontal, Positioner: true},
		Right:     &ResizeBar{Direction: Horizontal, Positioner: false},
		Thickness: unit.Dp(4),
	}
}

func (r Resizer) Layout(gtx layout.Context, dims layout.Dimensions) layout.Dimensions {
	defer op.Offset(r.position.Round()).Push(gtx.Ops).Pop()

	var off op.TransformStack

	off = op.Offset(image.Point{X: 0, Y: -gtx.Dp(r.Thickness)}).Push(gtx.Ops)
	r.Top.Layout(gtx, dims.Size.X, gtx.Dp(r.Thickness))
	off.Pop()

	off = op.Offset(image.Point{X: -gtx.Dp(unit.Dp(r.Thickness)), Y: 0}).Push(gtx.Ops)
	r.Left.Layout(gtx, gtx.Dp(r.Thickness), dims.Size.Y)
	off.Pop()

	off = op.Offset(image.Point{X: 0, Y: dims.Size.Y}).Push(gtx.Ops)
	r.Bottom.Layout(gtx, dims.Size.X, gtx.Dp(r.Thickness))
	off.Pop()

	off = op.Offset(image.Point{X: dims.Size.X, Y: 0}).Push(gtx.Ops)
	r.Right.Layout(gtx, gtx.Dp(r.Thickness), dims.Size.Y)
	off.Pop()

	return layout.Dimensions{}
}

func (r Resizer) Dragging() (bool, *ResizeBar) {
	if r.Top.Drag.Dragging() {
		return true, r.Top
	}
	if r.Bottom.Drag.Dragging() {
		return true, r.Bottom
	}
	if r.Left.Drag.Dragging() {
		return true, r.Left
	}
	if r.Right.Drag.Dragging() {
		return true, r.Right
	}
	return false, nil
}

func (r Resizer) Offset() f32.Point {
	if r.Top.Drag.Dragging() {
		return r.Top.Offset
	}

	if r.Bottom.Drag.Dragging() {
		return r.Bottom.Offset
	}

	if r.Left.Drag.Dragging() {
		return r.Left.Offset
	}

	if r.Right.Drag.Dragging() {
		return r.Right.Offset
	}

	return f32.Point{}
}

type ResizeBar struct {
	Direction  ResizeDirection
	Positioner bool

	StartPosition f32.Point

	Offset f32.Point

	Drag gesture.Drag
}

type ResizeDirection int

const (
	Vertical ResizeDirection = iota
	Horizontal
)

func (r *ResizeBar) Layout(gtx layout.Context, width int, height int) layout.Dimensions {
	for _, e := range r.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			r.StartPosition = e.Position
		case pointer.Drag:
			r.Offset = e.Position.Sub(r.StartPosition)
		case pointer.Release:
			r.Offset = f32.Point{}
			//ws.Window.Position = ws.Window.Position.Add(ws.Window.BottomBar.DragOffset)
		}
	}

	defer clip.Rect{Max: image.Pt(width, height)}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color.NRGBA{R: 255, A: 255})
	r.Drag.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(width, height)}
}
