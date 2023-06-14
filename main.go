package main

import (
	"log"
	"os"
	"ui/component"
	"ui/widget"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
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

	example := component.Example{
		List:  widget.List{List: layout.List{Axis: layout.Vertical}},
		Combo: widget.MakeCombo([]string{"one", "two"}, "numbers"),
	}
	win := widget.NewWindow("XQuery 1", f32.Pt(30, 30), example.Layout)

	example2 := component.Example{
		List:  widget.List{},
		Combo: widget.MakeCombo([]string{"one", "two", "three"}, "numbers"),
	}
	win2 := widget.NewWindow("XQuery 2", f32.Pt(40, 40), example2.Layout)

	//example3 := component.Example{}
	//win3 := widget.NewWindow("XQuery 3", f32.Pt(80, 80), example3.Layout)

	canvas := component.Canvas{
		Windows:  []*widget.Window{win, win2},
		GridSize: 0,
	}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			//paint.Fill(gtx.Ops, theme.Theme.Bg)

			//theme.Window(win).Layout(gtx)

			//op.Offset(image.Pt(300, 300)).Add(gtx.Ops)

			canvas.Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
