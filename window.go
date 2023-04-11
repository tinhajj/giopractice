package main

import (
	"image"
	"image/color"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
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
		defer clip.RRect{
			Rect: image.Rectangle{
				Min: image.Point{},
				Max: image.Pt(500, 500),
			},
			SE: 2,
			SW: 2,
			NW: 2,
			NE: 2,
		}.Push(ops).Pop()

		widge(gtx)
		return layout.Dimensions{Size: image.Pt(500, 500)}
	})
}
