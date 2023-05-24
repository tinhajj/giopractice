package widget

import (
	"image"

	"gioui.org/gesture"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
)

type ResizeBar struct {
	Width  unit.Dp
	Height unit.Dp

	Direction  ResizeDirection
	Positioner bool

	StartPosition float32

	Dragging bool
	Offset   float32

	Released bool

	Drag gesture.Drag
}

type ResizeDirection int

const (
	Vertical ResizeDirection = iota
	Horizontal
)

func (r *ResizeBar) Layout(gtx layout.Context) layout.Dimensions {
	r.Released = false

	for range r.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
	}

	defer clip.Rect{Max: image.Pt(gtx.Dp(r.Width), gtx.Dp(r.Height))}.Push(gtx.Ops).Pop()
	r.Drag.Add(gtx.Ops)

	return layout.Dimensions{
		Size: image.Point{X: gtx.Dp(r.Width), Y: gtx.Dp(r.Height)},
	}
}
