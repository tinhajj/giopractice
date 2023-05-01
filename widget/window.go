package widget

import "image"

type Window struct {
	Height   int
	Width    int
	Position image.Point
	Dragging bool
}
