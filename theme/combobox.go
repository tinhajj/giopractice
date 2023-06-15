package theme

import (
	"fmt"
	"ui/widget"

	"gioui.org/layout"
)

// ComboStyle holds combobox rendering parameters
type ComboStyle struct {
	w *widget.Combo
}

// Combo constructs c ComboStyle
func Combo(w *widget.Combo) ComboStyle {
	return ComboStyle{
		w: w,
	}
}

// Layout a combobox
func (c ComboStyle) Layout(gtx layout.Context) layout.Dimensions {
	return c.w.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		text := fmt.Sprintf("<%s>", c.w.Hint())

		subwidgets := make([]ButtonStyle, 0)
		subwidgets = append(subwidgets, Button(c.w.SelectButton(), text))

		if c.w.IsExpanded() {
			N := c.w.Len()
			for i := 0; i < N; i++ {
				bs := Button(c.w.Button(i), c.w.Item(i))
				subwidgets = append(subwidgets, bs)
			}
		}

		rigids := make([]layout.FlexChild, 0)
		for _, bs := range subwidgets {
			bs := bs
			rigids = append(rigids, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return bs.Layout(gtx)
			}))
		}

		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rigids...)
	})
}
