package component

import (
	"ui/theme"
	"ui/widget"

	"gioui.org/layout"
)

type Example struct {
	List widget.List

	submitBtn widget.Clickable
	cancelBtn widget.Clickable

	lightsCheckbox widget.Bool
	bigCheckbox    widget.Bool
	flipCheckbox   widget.Bool
}

func NewExample() Example {
	return Example{}
}

func (e *Example) Layout(gtx layout.Context) layout.Dimensions {
	//th := material.NewTheme(gofont.Collection())

	return theme.List(&e.List).Layout(gtx, 9, func(gtx layout.Context, i int) layout.Dimensions {
		if i == 0 {
			dim := theme.Button(&e.submitBtn, "Submit").Layout(gtx)
			return dim
		} else if i == 1 {
			dim := layout.Spacer{Height: 500, Width: 80}.Layout(gtx)
			return dim
		} else if i == 2 {
			dim := theme.Checkbox(&e.lightsCheckbox, "Lights").Layout(gtx)
			return dim
		} else if i == 3 {
			dim := layout.Spacer{Width: 20}.Layout(gtx)
			return dim
		} else if i == 4 {
			dim := theme.Checkbox(&e.bigCheckbox, "Big").Layout(gtx)
			return dim
		} else if i == 5 {
			dim := layout.Spacer{Width: 40}.Layout(gtx)
			return dim
		} else if i == 6 {
			dim := theme.Checkbox(&e.flipCheckbox, "Flip").Layout(gtx)
			return dim
		} else if i == 7 {
			dim := layout.Spacer{Width: 60}.Layout(gtx)
			return dim
		} else {
			dim := theme.Button(&e.cancelBtn, "Cancel").Layout(gtx)
			return dim
		}
	})
}
