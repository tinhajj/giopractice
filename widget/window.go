package widget

import (
	"gioui.org/f32"
	"gioui.org/gesture"
)

type Window struct {
	Height   int
	Width    int
	Position f32.Point
	Dragging bool

	Drag gesture.Drag

	StartClickPosition f32.Point
	StartPosition      f32.Point
	LastPosition       f32.Point
}
