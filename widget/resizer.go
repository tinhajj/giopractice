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

type Resizer struct {
	Top    *ResizeBar
	Bottom *ResizeBar
	Left   *ResizeBar
	Right  *ResizeBar

	Window *Window

	Thickness unit.Dp

	start  f32.Point
	offset f32.Point
}

func NewResizer(win *Window) *Resizer {
	return &Resizer{
		Top:       &ResizeBar{Direction: Vertical, Positioner: true, Window: win, Kind: Opposite},
		Bottom:    &ResizeBar{Direction: Vertical, Positioner: false, Window: win, Kind: Normal},
		Left:      &ResizeBar{Direction: Horizontal, Positioner: true, Window: win, Kind: Opposite},
		Right:     &ResizeBar{Direction: Horizontal, Positioner: false, Window: win, Kind: Normal},
		Window:    win,
		Thickness: unit.Dp(4),
	}
}

func (r Resizer) Layout(gtx layout.Context) layout.Dimensions {
	if ok, b := r.Dragging(); ok {
		defer op.Offset(b.StartWindow.Position.Round()).Push(gtx.Ops).Pop()
	} else {
		defer op.Offset(r.Window.Position.Round()).Push(gtx.Ops).Pop()
	}

	var off op.TransformStack

	var dims layout.Dimensions
	if ok, b := r.Dragging(); ok {
		win := b.StartWindow
		dims = layout.Dimensions{
			Size:     image.Point{X: gtx.Dp(win.Width), Y: gtx.Dp(win.Height)},
			Baseline: 0,
		}
	} else {
		win := r.Window
		dims = layout.Dimensions{
			Size:     image.Point{X: gtx.Dp(win.Width), Y: gtx.Dp(win.Height)},
			Baseline: 0,
		}
	}

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

type ResizeBar struct {
	Direction  ResizeDirection
	Positioner bool
	Window     *Window

	Kind Kind

	StartWindow       Window
	StartDragPosition f32.Point

	Offset f32.Point

	Drag gesture.Drag
}

type ResizeDirection int

type Kind int

const (
	Normal   Kind = 1
	Opposite Kind = -1
)

const (
	Vertical ResizeDirection = iota
	Horizontal
)

func (r *ResizeBar) Layout(gtx layout.Context, width int, height int) layout.Dimensions {
	for _, e := range r.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		switch e.Type {
		case pointer.Press:
			r.StartDragPosition = e.Position
			r.StartWindow = *r.Window
		case pointer.Drag:
			r.Offset = e.Position.Sub(r.StartDragPosition)

			if r.Direction == Vertical {
				pxHeight := gtx.Dp(r.StartWindow.Height) + (r.Offset.Round().Y * int(r.Kind))
				r.Window.Height = gtx.Metric.PxToDp(pxHeight)

				if r.Positioner {
					r.Window.Position.Y = r.StartWindow.Position.Add(r.Offset).Y
				}
			}

			if r.Direction == Horizontal {
				pxWidth := gtx.Dp(r.StartWindow.Width) + (r.Offset.Round().X * int(r.Kind))
				r.Window.Width = gtx.Metric.PxToDp(pxWidth)

				if r.Positioner {
					r.Window.Position.X = r.StartWindow.Position.Add(r.Offset).X
				}
			}
		case pointer.Release:
			//r.Offset = f32.Point{}
			//ws.Window.Position = ws.Window.Position.Add(ws.Window.BottomBar.DragOffset)
		}
	}

	defer clip.Rect{Max: image.Pt(width, height)}.Push(gtx.Ops).Pop()
	r.Drag.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(width, height)}
}
