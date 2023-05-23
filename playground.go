package main

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// Test colors.
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 150}
	red        = color.NRGBA{R: 255, G: 0, B: 0, A: 150}
	green      = color.NRGBA{R: 0, G: 255, B: 0, A: 150}
	blue       = color.NRGBA{R: 0, G: 0, B: 255, A: 150}
	purple     = color.NRGBA{R: 255, G: 0, B: 255, A: 150}
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	fmt.Printf("color: %+v size: %+v\n", color, size)
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func stacked(gtx layout.Context) layout.Dimensions {
	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		// Force widget to the same size as the second.
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// This will have a minimum constraint of 100x100.
			return ColorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 30), green)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(30, 100), blue)
		}),
	)
}
