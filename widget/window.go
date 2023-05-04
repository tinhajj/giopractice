package widget

import (
	"gioui.org/f32"
)

type Window struct {
	Height   int
	Width    int
	Position f32.Point
	Dragging bool

	LastPosition f32.Point
}
