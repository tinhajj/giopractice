package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"ui/theme"
	"ui/widget"

	tlayout "ui/layout"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
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
	th := material.NewTheme(gofont.Collection())

	active := true
	b := widget.NewBool(&active)
	clickable := widget.Clickable{}

	for windowEvent := range w.Events() {
		switch e := windowEvent.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			paint.Fill(gtx.Ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 100})

			cs := theme.Checkbox(b)
			lay := tlayout.ConstrainedLayout{Size: image.Pt(30, 30)}

			layout.UniformInset(unit.Dp(12)).Layout(gtx, func(gtx C) D {
				//return cs.Layout(gtx)
				return lay.Layout(gtx, cs.Layout)
			})

			//dims := material.Button(th, &clickable, "info").Layout(gtx)

			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// We insert two rigid elements:
				// First a button ...
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &clickable, "First")
						return btn.Layout(gtx)
					},
				),
				// ... then an empty spacer
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
				// We insert two rigid elements:
				// First a button ...
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &clickable, "Second")
						return btn.Layout(gtx)
					},
				),
				// ... then an empty spacer
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
