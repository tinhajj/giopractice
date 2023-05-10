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
			app.Size(unit.Dp(600), unit.Dp(600)),
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
	// th := material.NewTheme(gofont.Collection())
	// c := widget.Clickable{}
	// window := widget.NewWindow("Main")

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			op.Offset(image.Point{X: 100, Y: 100}).Add(gtx.Ops)
			//theme.Window(window).Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
