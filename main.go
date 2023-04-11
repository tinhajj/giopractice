package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
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
		width:     0,
		height:    510,
	}

	window := Window{tag: new(bool)}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			op.Offset(image.Point{X: 10, Y: 10}).Add(&ops)

			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			bv1.Layout(gtx)
			op.Offset(image.Point{X: 0, Y: 200}).Add(&ops)
			//bv2.Layout(gtx)
			op.Offset(image.Point{X: 0, Y: 200}).Add(&ops)
			window.Layout(gtx, bv2.Layout)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
