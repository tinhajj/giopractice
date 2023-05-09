package widget

import (
	"gioui.org/f32"
	"gioui.org/gesture"
)

type Window struct {
	Title    string
	Height   int
	Width    int
	Position f32.Point

	TopBar    *ResizeBar
	BottomBar *ResizeBar
	LeftBar   *ResizeBar
	RightBar  *ResizeBar
}

func NewWindow(title string) *Window {
	height := 800
	width := 400

	return &Window{
		Title:    title,
		Height:   height,
		Width:    width,
		Position: f32.Point{},

		TopBar:    NewResizeBar(15, width),
		BottomBar: NewResizeBar(15, width),
		LeftBar:   NewResizeBar(height, 15),
		RightBar:  NewResizeBar(height, 15),
	}
}

func (w *Window) Dragging() bool {
	return w.TopBar.Dragging || w.BottomBar.Dragging || w.LeftBar.Dragging || w.RightBar.Dragging
}

func (w *Window) DragOffset() f32.Point {
	if w.TopBar.Dragging {
		return w.TopBar.DragOffset
	}
	if w.BottomBar.Dragging {
		return w.BottomBar.DragOffset
	}
	if w.LeftBar.Dragging {
		return w.LeftBar.DragOffset
	}
	if w.RightBar.Dragging {
		return w.RightBar.DragOffset
	}

	return f32.Point{}
}

type ResizeBar struct {
	Height int
	Width  int

	Dragging bool

	Drag gesture.Drag

	StartPosition f32.Point
	DragOffset    f32.Point
}

func NewResizeBar(height, width int) *ResizeBar {
	return &ResizeBar{
		Height:        height,
		Width:         width,
		Dragging:      false,
		Drag:          gesture.Drag{},
		StartPosition: f32.Point{},
		DragOffset:    f32.Point{},
	}
}
