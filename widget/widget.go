package widget

import (
	"gioui.org/io/pointer"
	"gioui.org/widget"
)

type Draggable = widget.Draggable
type Border = widget.Border
type Clickable = widget.Clickable
type Press = widget.Press
type Label = widget.Label
type Bool = widget.Bool
type List = widget.List
type Scrollbar = widget.Scrollbar

var AllPointers = pointer.Press | pointer.Release | pointer.Move | pointer.Drag | pointer.Enter | pointer.Leave | pointer.Scroll
