package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/pointer"
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

	var outerClick bool
	var innerClick bool

	//var startClick f32.Point
	// th := material.NewTheme(gofont.Collection())
	// c := widget.Clickable{}
	// window := widget.NewWindow("Main")

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			for _, e := range gtx.Events(&outerClick) {
				if x, ok := e.(pointer.Event); ok {
					switch x.Type {
					case pointer.Press:
						fmt.Println("outer click")
					}
				}
			}

			for _, e := range gtx.Events(&innerClick) {
				if x, ok := e.(pointer.Event); ok {
					switch x.Type {
					case pointer.Press:
						fmt.Println("inner click")
					}
				}
			}

			stack1 := clip.Rect{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: gtx.Constraints.Max.X, Y: gtx.Constraints.Max.Y},
			}.Push(gtx.Ops)
			//pointer.PassOp{}.Push(gtx.Ops)
			pointer.InputOp{
				Tag:          &outerClick,
				Grab:         false,
				Types:        pointer.Press,
				ScrollBounds: image.Rectangle{},
			}.Add(gtx.Ops)
			paint.Fill(gtx.Ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})
			stack1.Pop()

			stack2 := clip.Rect{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 100, Y: 100},
			}.Push(gtx.Ops)
			pointer.InputOp{
				Tag:          &innerClick,
				Grab:         false,
				Types:        pointer.Press,
				ScrollBounds: image.Rectangle{},
			}.Add(gtx.Ops)
			paint.Fill(gtx.Ops, color.NRGBA{R: 100, G: 200, B: 0, A: 255})
			stack2.Pop()

			//theme.Window(window).Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
