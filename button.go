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
)

type ButtonVisual struct {
	roundness int
	pressed   bool
	tag       *bool
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
	// Process events that arrived between the last frame and this one.
	for _, ev := range q.Events(b.tag) {
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
				b.pressed = true
			case pointer.Release:
				b.pressed = false
			}
		}
	}

	// Confine the area of interest to a 100x100 rectangle.
	defer clip.RRect{
		Rect: image.Rectangle{Max: image.Pt(100, 100)},
		SE:   b.roundness,
		SW:   b.roundness,
		NW:   b.roundness,
		NE:   b.roundness,
	}.Push(ops).Pop()

	// Declare the tag.
	pointer.InputOp{
		Tag:          b.tag,
		Grab:         false,
		Types:        pointer.Press | pointer.Release,
		ScrollBounds: image.Rectangle{},
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
