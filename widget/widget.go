package widget

import (
	"ui/layout"

	"gioui.org/io/pointer"
	"gioui.org/widget"
)

type Draggable = widget.Draggable
type Border = widget.Border
type Clickable = widget.Clickable
type Press = widget.Press
type Label = widget.Label
type Bool = widget.Bool
type Editor = widget.Editor

var AllPointers = pointer.Press | pointer.Release | pointer.Move | pointer.Drag | pointer.Enter | pointer.Leave | pointer.Scroll

func Rigid(widget layout.Widget) layout.FlexChild {
	return layout.Rigid(widget)
}

func Rigids(widgets ...layout.Widget) []layout.FlexChild {
	children := []layout.FlexChild{}

	for _, w := range widgets {
		children = append(children, Rigid(w))
	}

	return children
}
