package theme

import (
	"fmt"
	"image"
	"ui/widget"

	"ui/layout"
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

		if c.w.HasSelected() {
			text = fmt.Sprintf("<%s>", c.w.SelectedText())
		}

		subwidgets := make([]ButtonStyle, 0)
		subwidgets = append(subwidgets, Button(c.w.SelectButton(), text))

		if c.w.IsExpanded() {
			N := c.w.Len()
			for i := 0; i < N; i++ {
				bs := Button(c.w.Button(i), c.w.Item(i))
				subwidgets = append(subwidgets, bs)
			}
		}

		funcs := []layout.Widget{}
		for _, w := range subwidgets {
			funcs = append(funcs, w.Layout)
		}
		_, dim := layout.Largest(gtx, funcs...)

		rigids := make([]layout.FlexChild, 0)
		for _, bs := range subwidgets {
			bs := bs
			ngtx := gtx
			ngtx.Constraints = layout.Exact(image.Pt(dim.Size.X, dim.Size.Y))
			rigids = append(rigids, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return bs.Layout(ngtx)
			}))
		}

		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx, rigids...)
	})
}
