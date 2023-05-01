package theme

import (
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
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
	// Process events that arrived between the last frame and this one.
	for _, ev := range gtx.Events(ws.Window) {
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Drag:
				ws.Window.Dragging = true
			case pointer.Release:
				ws.Window.Dragging = false
			}
		}
	}

	op.Offset(ws.Window.Position).Add(gtx.Ops)

	rect := clip.Rect{
		Min: image.Point{
			X: 0,
			Y: 5,
		},
		Max: ws.Window.Position.Add(image.Point{X: 0, Y: 10}),
	}
	area := rect.Push(gtx.Ops)
	pointer.InputOp{
		Tag:          ws.Window,
		Grab:         false,
		Types:        pointer.Drag | pointer.Release,
		ScrollBounds: image.Rectangle{},
	}.Add(gtx.Ops)
	if ws.Window.Dragging {
		paint.Fill(gtx.Ops, color.NRGBA{100, 255, 0, 255})
		pointer.CursorNorthResize.Add(gtx.Ops)
	} else {
		paint.Fill(gtx.Ops, color.NRGBA{100, 0, 255, 255})
	}
	pointer.CursorNorthResize.Add(gtx.Ops)
	area.Pop()

	op.Offset(image.Point{0, 10}).Add(gtx.Ops)

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

		defer clip.Stroke{
			Path:  cRect.Path(gtx.Ops),
			Width: 10,
		}.Op().Push(gtx.Ops).Pop()
		//paint.Fill(gtx.Ops, color.NRGBA{100, 155, 255, 255})

		return layout.Dimensions{
			Size:     gtx.Constraints.Max,
			Baseline: 0,
		}
	})
}
