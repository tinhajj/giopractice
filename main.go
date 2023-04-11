package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Canvas"),
			app.Size(unit.Dp(650), unit.Dp(600)),
		)

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops
	bv1 := &ButtonVisual{
		roundness: 2,
		pressed:   false,
		tag:       new(bool),
		width:     0,
		height:    0,
	}

	bv2 := &ButtonVisual{
		roundness: 0,
		pressed:   false,
		tag:       new(bool),
		width:     300,
		height:    510,
	}

	bv3 := &ButtonVisual{
		roundness: 0,
		pressed:   false,
		tag:       new(bool),
		width:     500,
		height:    510,
	}

	th := material.NewTheme(gofont.Collection())

	window := Window{
		tag: new(bool),
		list: widget.List{Scrollbar: widget.Scrollbar{}, List: layout.List{
			Axis: layout.Vertical,
		}},
		theme:   th,
		widgets: []layout.Widget{bv2.Layout, bv3.Layout},
	}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			op.Offset(image.Point{X: 10, Y: 10}).Add(&ops)

			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			bv1.Layout(gtx)
			op.Offset(image.Point{X: 0, Y: 200}).Add(&ops)
			window.Layout(gtx)

			op.Offset(image.Point{X: 0, Y: 200}).Add(&ops)

			bv3.Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
