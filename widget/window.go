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
	height := 300
	width := 400

	return &Window{
		Title:    title,
		Height:   300,
		Width:    400,
		Position: f32.Point{},

		TopBar:    NewResizeBar(15, width),
		BottomBar: NewResizeBar(15, width),
		LeftBar:   NewResizeBar(height, 15),
		RightBar:  NewResizeBar(height, 15),
	}
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
