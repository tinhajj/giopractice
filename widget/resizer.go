package widget

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/gesture"
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

func (r Resizer) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	dims := w(gtx)

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

	return dims
}

type ResizeBar struct {
	Direction  ResizeDirection
	Positioner bool

	StartPosition float32

	Dragging bool
	Offset   float32

	Drag gesture.Drag
}

type ResizeDirection int

const (
	Vertical ResizeDirection = iota
	Horizontal
)

func (r *ResizeBar) Layout(gtx layout.Context, width int, height int) layout.Dimensions {
	for _, e := range r.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		fmt.Println(e)
	}

	defer clip.Rect{Max: image.Pt(width, height)}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color.NRGBA{R: 255, A: 255})
	r.Drag.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(width, height)}
}
