package main

import (
	"log"
	"os"
	"ui/component"
	"ui/theme"
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

	//active := true
	//b := widget.NewBool(&active)

	example := component.Example{}
	win := widget.NewWindow("XQuery", f32.Pt(100, 100), example.Layout)

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			paint.Fill(gtx.Ops, theme.Theme.Bg)
			theme.Window(win).Layout(gtx)

			//op.Offset(image.Pt(300, 300)).Add(gtx.Ops)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
