package main

import (
	"image"
	"image/color"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Window struct {
	tag     *bool
	list    widget.List
	theme   *material.Theme
	widgets []layout.Widget
}

func (w *Window) Layout(gtx layout.Context) layout.Dimensions {
	return w.draw(gtx.Ops, gtx.Queue, gtx)
}

func (w *Window) draw(ops *op.Ops, q event.Queue, gtx layout.Context) layout.Dimensions {
	// Process events that arrived between the last frame and this one.
	for _, ev := range q.Events(w.tag) {
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
			case pointer.Release:
			}
		}
	}

	gtx.Constraints.Min = image.Point{X: 0, Y: 0}
	gtx.Constraints.Max = image.Point{X: 100, Y: 199}

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
		//defer clip.RRect{
		//	Rect: image.Rectangle{
		//		Min: gtx.Constraints.Min,
		//		Max: gtx.Constraints.Max,
		//	},
		//	SE: 2,
		//	SW: 2,
		//	NW: 2,
		//	NE: 2,
		//}.Push(ops).Pop()

		/*
			style := material.ScrollbarStyle{
				Scrollbar: &w.list.Scrollbar,
				Track: material.ScrollTrackStyle{
					MajorPadding: 0,
					MinorPadding: 0,
					Color:        color.NRGBA{},
				},
				Indicator: material.ScrollIndicatorStyle{},
			}
		*/

		return material.List(w.theme, &w.list).Layout(gtx, len(w.widgets), func(gtx layout.Context, index int) layout.Dimensions {
			return w.widgets[index](gtx)
		})

		return layout.Dimensions{
			Size:     image.Pt(300, 300),
			Baseline: 0,
		}
	})
}
