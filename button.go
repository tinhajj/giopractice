package main

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type ButtonVisual struct {
	pressed bool
	tag     bool
}

func (b *ButtonVisual) Layout(gtx layout.Context) layout.Dimensions {
	col := color.NRGBA{
		R: 0x80,
		G: 0,
		B: 0,
		A: 0xFF,
	}
	return b.drawSquare(gtx.Ops, gtx.Queue, col)
}

func (b *ButtonVisual) drawSquare(ops *op.Ops, q event.Queue, baseColor color.NRGBA) layout.Dimensions {
	fmt.Println("drawsquare")
	// Process events that arrived between the last frame and this one.
	for _, ev := range q.Events(0) {
		fmt.Println("loop")
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
				fmt.Println("press")
				b.pressed = true
			case pointer.Release:
				fmt.Println("release")
				b.pressed = false
			}
		}
	}

	// Confine the area of interest to a 100x100 rectangle.
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()

	// Declare the tag.
	pointer.InputOp{
		Tag:   0,
		Types: pointer.Press | pointer.Release,
	}.Add(ops)

	var c color.NRGBA
	if b.pressed {
		c = color.NRGBA{R: 0xFF, A: 0xFF}
	} else {
		c = baseColor
	}
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)
	return layout.Dimensions{Size: image.Pt(100, 100)}
}
