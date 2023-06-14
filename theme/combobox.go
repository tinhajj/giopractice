package theme

import (
	"fmt"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/unit"
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
	if c.w.SelectButton().Clicked() {
		c.w.Toggle()
	}

	//fmt.Println(c.w.IsExpanded())

	for i := 0; i < c.w.Len(); i++ {
		if c.w.Button(i).Clicked() {
			if err := c.w.SelectIndex(i); err != nil {
				fmt.Println("giox error: bad index")
			}
			c.w.Toggle()
		}
	}

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

	var inset float32 = 0.0
	if c.w.IsExpanded() {
		inset = 10
	}

	rigids := make([]layout.FlexChild, 0)
	for _, bs := range subwidgets {
		bs := bs
		rigids = append(rigids, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return bs.Layout(gtx)
		}))
	}

	return layout.Inset{Left: unit.Dp(inset)}.Layout(gtx, func(layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rigids...)
	})
}
