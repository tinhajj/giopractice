package main

import (
	"image/color"
	"log"
	"os"
	"ui/theme"
	"ui/widget"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
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

	fonts := gofont.Collection()
	shaper := text.NewShaper(fonts)

	font := text.Font{}

	//active := true
	//b := widget.NewBool(&active)

	win := widget.NewWindow("Hello")
	label := widget.Label{Alignment: text.Middle}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			paint.Fill(gtx.Ops, color.NRGBA{R: 105, G: 105, B: 105, A: 255})

			textColorMacro := op.Record(gtx.Ops)
			paint.ColorOp{Color: color.NRGBA{R: 0, B: 0, G: 0, A: 255}}.Add(gtx.Ops)
			textColor := textColorMacro.Stop()
			label.Layout(gtx, shaper, font, unit.Sp(30), "Hello", textColor)

			theme.Window(win).Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
