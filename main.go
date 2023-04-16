package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"ui/theme"

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
	th := material.NewTheme(gofont.Collection())
	c := widget.Clickable{}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			op.Offset(image.Point{X: 10, Y: 10}).Add(&ops)

			gtx.Constraints = layout.Constraints{
				Min: image.Point{
					X: 100,
					Y: 100,
				},
				Max: image.Point{
					X: 600,
					Y: 300,
				},
			}

			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			theme.Button(th, &c, "Click me, click me").Layout(gtx)
			op.Offset(image.Point{X: 0, Y: 200}).Add(&ops)

			layout.Stack{}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return theme.Button(th, &c, "Click me, click me").Layout(gtx)
				}),
			)

			op.Offset(image.Point{X: 100, Y: 500}).Add(&ops)

			material.Button(th, &c, "Click me, click me").Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
