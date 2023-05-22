package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"ui/widget"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type C = layout.Context
type D = layout.Dimensions

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
	//th := material.NewTheme(gofont.Collection())

	//active := true
	//b := widget.NewBool(&active)

	win := widget.NewWindow("XQuery")
	win.Position = f32.Point{X: 100, Y: 100}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			paint.Fill(gtx.Ops, color.NRGBA{R: 105, G: 105, B: 105, A: 255})

			//theme.Window(win).Layout(gtx)

			op.Offset(image.Pt(300, 300)).Add(gtx.Ops)
			stacked(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
