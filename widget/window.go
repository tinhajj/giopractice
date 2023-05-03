package widget

import (
	"image"

	"gioui.org/f32"
)

type Window struct {
	Height   int
	Width    int
	Position image.Point
	Dragging bool

	LastPosition f32.Point
}
