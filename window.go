package main

import (
	"image"
	"image/color"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

type Window struct {
	tag *bool
}

func (w *Window) Layout(gtx layout.Context, widge layout.Widget) layout.Dimensions {
	return w.draw(gtx.Ops, gtx.Queue, gtx, widge)
}

func (w *Window) draw(ops *op.Ops, q event.Queue, gtx layout.Context, widge layout.Widget) layout.Dimensions {
	// Process events that arrived between the last frame and this one.
	for _, ev := range q.Events(w.tag) {
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
			case pointer.Release:
			}
		}
	}

	border := widget.Border{
		Color: color.NRGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
		CornerRadius: 1,
		Width:        1,
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// Confine the area of interest to a 100x100 rectangle.
		defer clip.RRect{
			Rect: image.Rectangle{
				Min: image.Point{},
				Max: image.Pt(100, 100),
			},
			SE: 2,
			SW: 2,
			NW: 2,
			NE: 2,
		}.Push(ops).Pop()

		paint.ColorOp{Color: color.NRGBA{G: 0xFF, A: 0xFF}}.Add(ops)
		paint.PaintOp{}.Add(ops)
		return layout.Dimensions{Size: image.Pt(500, 500)}
	})
}
