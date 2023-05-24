package theme

import (
	"image"
	"image/color"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
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
	defer op.Offset(ws.Window.Position.Round()).Push(gtx.Ops).Pop()

	return widget.OuterBorder{
		Border: widget.Border{
			Color:        color.NRGBA{A: 255, R: 85, G: 170, B: 170},
			CornerRadius: 0,
			Width:        unit.Dp(1),
		},
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ws.Window.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			height := gtx.Dp(ws.Window.Height)
			width := gtx.Dp(ws.Window.Width)
			gtx.Constraints.Min = image.Point{}
			gtx.Constraints.Max = image.Point{X: width, Y: height}

			defer clip.Rect{Max: image.Point{X: width, Y: height}}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, Yellow)

			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					bg := widget.Background{Color: color.NRGBA{R: 234, G: 255, B: 255, A: 255}}
					return bg.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return Label(gtx, unit.Sp(20), ws.Window.Title)
						})
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return widget.HR{Width: unit.Dp(1), Color: Black}.Layout(gtx)
				}),
			)

			return layout.Dimensions{
				Size: image.Point{X: width, Y: height},
			}
		})
	})

	// Process events that arrived between the last frame and this one.
	//for _, e := range ws.Window.TopBar.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
	//	switch e.Type {
	//	case pointer.Press:
	//		ws.Window.TopBar.StartPosition = e.Position
	//		ws.Window.TopBar.StartWindowHeight = ws.Window.Height
	//	case pointer.Drag:
	//		ws.Window.TopBar.Dragging = true
	//		ws.Window.TopBar.DragOffset = e.Position.Sub(ws.Window.TopBar.StartPosition)
	//		ws.Window.Height = ws.Window.TopBar.StartWindowHeight - ws.Window.TopBar.DragOffset.Round().Y
	//	case pointer.Release:
	//		ws.Window.TopBar.Dragging = false
	//		ws.Window.Position = ws.Window.Position.Add(ws.Window.TopBar.DragOffset)
	//	}
	//}

	//for _, e := range ws.Window.BottomBar.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
	//	switch e.Type {
	//	case pointer.Press:
	//		ws.Window.BottomBar.StartPosition = e.Position
	//	case pointer.Drag:
	//		ws.Window.BottomBar.Dragging = true
	//		ws.Window.BottomBar.DragOffset = e.Position.Sub(ws.Window.BottomBar.StartPosition)
	//	case pointer.Release:
	//		ws.Window.BottomBar.Dragging = false
	//		ws.Window.Position = ws.Window.Position.Add(ws.Window.BottomBar.DragOffset)
	//	}
	//}

	//for _, e := range ws.Window.LeftBar.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
	//	switch e.Type {
	//	case pointer.Press:
	//		ws.Window.LeftBar.StartPosition = e.Position
	//	case pointer.Drag:
	//		ws.Window.LeftBar.Dragging = true
	//		ws.Window.LeftBar.DragOffset = e.Position.Sub(ws.Window.LeftBar.StartPosition)
	//	case pointer.Release:
	//		ws.Window.LeftBar.Dragging = false
	//		ws.Window.Position = ws.Window.Position.Add(ws.Window.LeftBar.DragOffset)
	//	}
	//}

	//for _, e := range ws.Window.RightBar.Drag.Events(gtx.Metric, gtx.Queue, gesture.Both) {
	//	switch e.Type {
	//	case pointer.Press:
	//		ws.Window.RightBar.StartPosition = e.Position
	//	case pointer.Drag:
	//		ws.Window.RightBar.Dragging = true
	//		ws.Window.RightBar.DragOffset = e.Position.Sub(ws.Window.RightBar.StartPosition)
	//	case pointer.Release:
	//		ws.Window.RightBar.Dragging = false
	//		ws.Window.Position = ws.Window.Position.Add(ws.Window.RightBar.DragOffset)
	//	}
	//}

	//op.Offset(ws.Window.Position.Round()).Push(gtx.Ops)

	//rect := clip.Rect{
	//	Min: image.Point{X: 0, Y: 0},
	//	Max: image.Point{X: ws.Window.TopBar.Width, Y: ws.Window.TopBar.Height},
	//}
	//stack := rect.Push(gtx.Ops)
	//pointer.CursorNorthResize.Add(gtx.Ops)
	//ws.Window.TopBar.Drag.Add(gtx.Ops)
	//stack.Pop()

	//rect = clip.Rect{
	//	Min: image.Point{X: 0, Y: ws.Window.Height - ws.Window.BottomBar.Height},
	//	Max: image.Point{X: ws.Window.Width, Y: ws.Window.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//pointer.CursorSouthResize.Add(gtx.Ops)
	//ws.Window.BottomBar.Drag.Add(gtx.Ops)
	//stack.Pop()

	//rect = clip.Rect{
	//	Min: image.Point{X: 0, Y: 0},
	//	Max: image.Point{X: ws.Window.LeftBar.Width, Y: ws.Window.LeftBar.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//ws.Window.LeftBar.Drag.Add(gtx.Ops)
	//pointer.CursorColResize.Add(gtx.Ops)
	//stack.Pop()

	//rect = clip.Rect{
	//	Min: image.Point{X: ws.Window.Width - ws.Window.RightBar.Width, Y: 0},
	//	Max: image.Point{X: ws.Window.Width, Y: ws.Window.RightBar.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//ws.Window.RightBar.Drag.Add(gtx.Ops)
	//pointer.CursorColResize.Add(gtx.Ops)
	//stack.Pop()

	//if ws.Window.Dragging() {
	//	op.Offset(ws.Window.DragOffset().Round()).Push(gtx.Ops)
	//}

	//rect = clip.Rect{
	//	Min: image.Point{X: 0, Y: 0},
	//	Max: image.Point{X: ws.Window.TopBar.Width, Y: ws.Window.TopBar.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	//stack.Pop()

	//rect = clip.Rect{
	//	Min: image.Point{X: 0, Y: ws.Window.Height - ws.Window.BottomBar.Height},
	//	Max: image.Point{X: ws.Window.Width, Y: ws.Window.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	//stack.Pop()

	//rect = clip.Rect{
	//	Min: image.Point{X: 0, Y: 0},
	//	Max: image.Point{X: ws.Window.LeftBar.Width, Y: ws.Window.LeftBar.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	//stack.Pop()

	//rect = clip.Rect{
	//	Min: image.Point{X: ws.Window.Width - ws.Window.RightBar.Width, Y: 0},
	//	Max: image.Point{X: ws.Window.Width, Y: ws.Window.RightBar.Height},
	//}
	//stack = rect.Push(gtx.Ops)
	//paint.Fill(gtx.Ops, color.NRGBA{A: 200, B: 100})
	//stack.Pop()

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
