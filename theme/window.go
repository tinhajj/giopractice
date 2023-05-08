package theme

import (
	"fmt"
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
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
	for _, e := range ws.Window.TopBar.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
		fmt.Println(e)
		//switch e.Type {
		//case pointer.Press:
		//	ws.Window.StartClickPosition = e.Position
		//	ws.Window.StartPosition = ws.Window.Position
		//case pointer.Drag:
		//	ws.Window.Dragging = true

		//	ws.Window.DragOffset = e.Position.Sub(ws.Window.StartClickPosition)
		//	//ws.Window.Position = ws.Window.StartPosition.Add(difference)
		//case pointer.Release:
		//	ws.Window.Dragging = false

		//	ws.Window.DragOffset = e.Position.Sub(ws.Window.StartClickPosition)
		//	//ws.Window.Position = ws.Window.StartPosition.Add(difference)
		//}
	}

	rect := clip.Rect{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: ws.Window.TopBar.Width, Y: ws.Window.TopBar.Height},
	}
	stack := rect.Push(gtx.Ops)
	paint.Fill(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 100, A: 200})
	pointer.CursorNorthResize.Add(gtx.Ops)
	stack.Pop()

	rect = clip.Rect{
		Min: image.Point{X: 0, Y: ws.Window.Height - ws.Window.BottomBar.Height},
		Max: image.Point{X: ws.Window.Width, Y: ws.Window.Height},
	}
	stack = rect.Push(gtx.Ops)
	paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	pointer.CursorSouthResize.Add(gtx.Ops)
	stack.Pop()

	rect = clip.Rect{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: ws.Window.LeftBar.Width, Y: ws.Window.LeftBar.Height},
	}
	stack = rect.Push(gtx.Ops)
	paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	pointer.CursorColResize.Add(gtx.Ops)
	stack.Pop()

	rect = clip.Rect{
		Min: image.Point{X: ws.Window.Width - ws.Window.RightBar.Width, Y: 0},
		Max: image.Point{X: ws.Window.Width, Y: ws.Window.RightBar.Height},
	}
	stack = rect.Push(gtx.Ops)
	paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	pointer.CursorColResize.Add(gtx.Ops)
	stack.Pop()

	return layout.Dimensions{}

	//op.Offset(ws.Window.Position.Round()).Add(gtx.Ops)

	//rect := clip.Rect{
	//	Min: image.Point{
	//		X: 0,
	//		Y: 0,
	//	},
	//	Max: image.Point{X: ws.Window.Width, Y: 10},
	//}

	//area := rect.Push(gtx.Ops)
	//{
	//	ws.Window.Drag.Add(gtx.Ops)
	//	if ws.Window.Dragging {
	//		paint.Fill(gtx.Ops, color.NRGBA{100, 255, 0, 255})
	//		pointer.CursorNorthResize.Add(gtx.Ops)
	//	} else {
	//		paint.Fill(gtx.Ops, color.NRGBA{100, 0, 255, 255})
	//	}
	//	pointer.CursorNorthResize.Add(gtx.Ops)
	//}
	//area.Pop()

	//op.Offset(image.Point{0, 10}).Add(gtx.Ops)

	//point := image.Point{
	//	Y: ws.Window.Height,
	//	X: ws.Window.Width,
	//}
	//gtx.Constraints.Max = point

	//border := widget.Border{
	//	Color:        color.NRGBA{155, 25, 155, 255},
	//	CornerRadius: 2,
	//	Width:        1,
	//}

	//w := func(gtx layout.Context) layout.Dimensions {
	//	cRect := clip.UniformRRect(image.Rectangle{
	//		Min: image.Point{},
	//		Max: gtx.Constraints.Max,
	//	}, 2)
	//	defer cRect.Push(gtx.Ops).Pop()
	//	paint.Fill(gtx.Ops, color.NRGBA{100, 255, 255, 255})

	//	return layout.Dimensions{
	//		Size:     gtx.Constraints.Max,
	//		Baseline: 0,
	//	}
	//}

	//if !ws.Window.Dragging {
	//	border.Layout(gtx, w)
	//}

	//op.Offset(ws.Window.DragOffset.Round()).Add(gtx.Ops)
	//return border.Layout(gtx, w)
}
