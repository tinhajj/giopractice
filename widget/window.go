package widget

import (
	"image"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/unit"
)

type Window struct {
	Title    string
	Height   unit.Dp
	Width    unit.Dp
	Position f32.Point
}

func NewWindow(title string) *Window {
	height := unit.Dp(500)
	width := unit.Dp(400)

	return &Window{
		Title:    title,
		Height:   height,
		Width:    width,
		Position: f32.Point{},
	}
}

func (w *Window) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	height := gtx.Dp(w.Height)
	width := gtx.Dp(w.Width)

	widget(gtx)

	return layout.Dimensions{
		Size:     image.Pt(height, width),
		Baseline: 0,
	}
}

//func (w *Window) Dragging() bool {
//	return w.TopBar.Dragging || w.BottomBar.Dragging || w.LeftBar.Dragging || w.RightBar.Dragging
//}
//
//func (w *Window) DragOffset() f32.Point {
//	if w.TopBar.Dragging {
//		return w.TopBar.DragOffset
//	}
//	if w.BottomBar.Dragging {
//		return w.BottomBar.DragOffset
//	}
//	if w.LeftBar.Dragging {
//		return w.LeftBar.DragOffset
//	}
//	if w.RightBar.Dragging {
//		return w.RightBar.DragOffset
//	}
//
//	return f32.Point{}
//}
//
//type ResizeBar struct {
//	Height int
//	Width  int
//
//	Dragging bool
//
//	Drag gesture.Drag
//
//	StartPosition f32.Point
//	DragOffset    f32.Point
//
//	StartWindowHeight int
//}
//
//func NewResizeBar(height, width int) *ResizeBar {
//	return &ResizeBar{
//		Height:        height,
//		Width:         width,
//		Dragging:      false,
//		Drag:          gesture.Drag{},
//		StartPosition: f32.Point{},
//		DragOffset:    f32.Point{},
//	}
//}
