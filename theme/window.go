package theme

import (
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	gwidget "gioui.org/widget"
)

type WindowStyle struct {
	Window *widget.Window
}

func Window(window *widget.Window) WindowStyle {
	return WindowStyle{
		Window: window,
	}
}

func (ws WindowStyle) Layout(gtx layout.Context) layout.Dimensions {
	point := image.Point{
		Y: ws.Window.Height,
		X: ws.Window.Width,
	}
	gtx.Constraints.Max = point

	border := gwidget.Border{
		Color:        color.NRGBA{155, 25, 155, 255},
		CornerRadius: 2,
		Width:        1,
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		cRect := clip.UniformRRect(image.Rectangle{
			Min: image.Point{},
			Max: gtx.Constraints.Max,
		}, 2)
		defer cRect.Push(gtx.Ops).Pop()

		paint.Fill(gtx.Ops, color.NRGBA{100, 255, 255, 255})
		return layout.Dimensions{
			Size:     gtx.Constraints.Max,
			Baseline: 0,
		}
	})
}
