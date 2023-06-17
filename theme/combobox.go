package theme

import (
	"fmt"
	"image"
	"ui/widget"

	"ui/layout"

	"gioui.org/font"
	"gioui.org/op"
	"gioui.org/op/paint"
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
	return c.w.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		text := fmt.Sprintf("<%s>", c.w.Hint())

		if c.w.HasSelected() {
			text = fmt.Sprintf("<%s>", c.w.SelectedText())
		}

		// components are the individual components that make up the
		// dropdown, we round them up to figure out the largest one
		components := make([]layout.Widget, 0)
		components = append(components, Button(c.w.SelectButton(), text).Layout)

		if c.w.IsExpanded() {
			N := c.w.Len()
			for i := 0; i < N; i++ {
				item := comboItem(gtx, c.w.Button(i), c.w.Item(i))
				components = append(components, item)
			}
		}

		funcs := []layout.Widget{}
		for _, w := range components {
			funcs = append(funcs, w)
		}
		_, dim := layout.Largest(gtx, funcs...)

		ngtx := gtx
		ngtx.Constraints = layout.Exact(image.Pt(dim.Size.X, dim.Size.Y))

		btn := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return Button(c.w.SelectButton(), text).Layout(ngtx)
		})

		opts := make([]layout.FlexChild, 0)
		if c.w.IsExpanded() {
			N := c.w.Len()
			for i := 0; i < N; i++ {
				item := comboItem(gtx, c.w.Button(i), c.w.Item(i))
				opts = append(opts, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return item(ngtx)
				}))
			}
		}

		optsBox := widget.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widget.Border{
				Color: Theme.Black,
				Width: 1,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx, opts...)
			})
		})

		everything := []layout.FlexChild{btn}
		if c.w.IsExpanded() {
			everything = append(everything, optsBox)
		}

		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx, everything...)
	})
}

func comboItem(gtx layout.Context, c *widget.Clickable, label string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		colMacro := op.Record(gtx.Ops)
		paint.ColorOp{Color: Theme.Palette.Black}.Add(gtx.Ops)
		colOp := colMacro.Stop()

		return c.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			paint.Fill(gtx.Ops, Theme.White)
			return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return widget.Label{}.Layout(gtx, Theme.Shaper, font.Font{}, Theme.TextSize, label, colOp)
			})
		})

	}
}
