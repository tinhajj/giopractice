package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"ui/theme"
	"ui/widget"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
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
	window := widget.Window{
		Height: 300,
		Width:  100,
		Position: f32.Point{
			X: 100,
			Y: 10,
		},
	}
	drag := &widget.Draggable{Type: "wig"}
	widget := func(gtx layout.Context) layout.Dimensions {
		sz := image.Pt(10, 100)
		defer clip.Rect{
			Min: image.Point{0, 0},
			Max: image.Point{100, 100},
		}.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, color.NRGBA{G: 220, A: 255})
		return layout.Dimensions{Size: sz}
	}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			theme.Window(&window).Layout(gtx)

			op.Offset(image.Point{X: 100, Y: 300}).Add(gtx.Ops)
			drag.Layout(gtx, widget, widget)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
