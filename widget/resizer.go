package widget

import (
	"image"

	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
)

type resizer struct {
	top    *resizebar
	bottom *resizebar
	left   *resizebar
	right  *resizebar

	window *Window

	previousDim layout.Dimensions

	thickness unit.Dp
}

func newResizer(win *Window) *resizer {
	return &resizer{
		top:    &resizebar{direction: vertical, positioner: true, window: win, kind: opposite},
		bottom: &resizebar{direction: vertical, positioner: false, window: win, kind: normal},
		left:   &resizebar{direction: horizontal, positioner: true, window: win, kind: opposite},
		right:  &resizebar{direction: horizontal, positioner: false, window: win, kind: normal},

		window:    win,
		thickness: unit.Dp(3),
	}
}

func (r *resizer) update(gtx layout.Context) {
	r.top.update(gtx)
	r.bottom.update(gtx)
	r.left.update(gtx)
	r.right.update(gtx)
}

func (r *resizer) layout(gtx layout.Context, dim layout.Dimensions) layout.Dimensions {
	if ok, b := r.dragging(); ok {
		defer op.Offset(b.startWindow.Position.Round()).Push(gtx.Ops).Pop()
		dim = r.previousDim
	} else {
		defer op.Offset(r.window.Position.Round()).Push(gtx.Ops).Pop()
		r.previousDim = dim
	}

	thickness := gtx.Dp(r.thickness)

	var off op.TransformStack

	off = op.Offset(image.Pt(0, -thickness)).Push(gtx.Ops)
	r.top.layout(gtx, dim.Size.X, thickness)
	off.Pop()

	off = op.Offset(image.Pt(-thickness, 0)).Push(gtx.Ops)
	r.left.layout(gtx, thickness, dim.Size.Y)
	off.Pop()

	off = op.Offset(image.Pt(0, dim.Size.Y)).Push(gtx.Ops)
	r.bottom.layout(gtx, dim.Size.X, thickness)
	off.Pop()

	off = op.Offset(image.Pt(dim.Size.X, 0)).Push(gtx.Ops)
	r.right.layout(gtx, thickness, dim.Size.Y)
	off.Pop()

	return layout.Dimensions{}
}

func (r *resizer) dragging() (bool, *resizebar) {
	if r.top.drag.Dragging() {
		return true, r.top
	}
	if r.bottom.drag.Dragging() {
		return true, r.bottom
	}
	if r.left.drag.Dragging() {
		return true, r.left
	}
	if r.right.drag.Dragging() {
		return true, r.right
	}

	return false, nil
}

type resizebar struct {
	direction  resizeDirection
	positioner bool
	window     *Window

	kind kind

	startWindow       Window
	startDragPosition f32.Point

	offset f32.Point

	drag gesture.Drag
}

func (r *resizebar) update(gtx layout.Context) {
	for _, e := range r.drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			r.startDragPosition = e.Position
			r.startWindow = *r.window
		case pointer.Drag:
			r.offset = e.Position.Sub(r.startDragPosition)

			if r.direction == vertical || r.direction == both {
				pxHeight := gtx.Dp(r.startWindow.Height) + (r.offset.Round().Y * int(r.kind))
				height, clamped := r.clamp(gtx.Metric.PxToDp(pxHeight))
				r.window.Height = height

				if r.positioner && !clamped {
					r.window.Position.Y = r.startWindow.Position.Add(r.offset).Y
				}
			}

			if r.direction == horizontal || r.direction == both {
				pxWidth := gtx.Dp(r.startWindow.Width) + (r.offset.Round().X * int(r.kind))
				width, clamped := r.clamp(gtx.Metric.PxToDp(pxWidth))
				r.window.Width = width

				if r.positioner && !clamped {
					r.window.Position.X = r.startWindow.Position.Add(r.offset).X
				}
			}
		case pointer.Release:
		}
	}
}

func (r *resizebar) layout(gtx layout.Context, width int, height int) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(width, height)}.Push(gtx.Ops).Pop()
	if r.direction == horizontal {
		pointer.CursorEastResize.Add(gtx.Ops)
	}
	if r.direction == vertical {
		pointer.CursorNorthResize.Add(gtx.Ops)
	}
	//paint.Fill(gtx.Ops, color.NRGBA{G: 255, A: 100})
	r.drag.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(width, height)}
}

// clamp restricts the height or width of a window and if it was clamped
func (r *resizebar) clamp(next unit.Dp) (unit.Dp, bool) {
	if next < unit.Dp(300) {
		return unit.Dp(300), true
	}

	return next, false
}

type resizeDirection int

const (
	vertical resizeDirection = iota
	horizontal
	both
)

type kind int

const (
	normal   kind = 1
	opposite kind = -1
)
