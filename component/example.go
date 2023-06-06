package component

import (
	"fmt"
	"ui/theme"
	"ui/widget"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Example struct {
	list widget.List

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
	th := material.NewTheme(gofont.Collection())

	dim := material.List(th, &e.list).Layout(gtx, 9, func(gtx layout.Context, i int) layout.Dimensions {
		if i == 0 {
			return theme.Button(&e.submitBtn, "Submit").Layout(gtx)
		} else if i == 1 {
			return layout.Spacer{Width: 20}.Layout(gtx)
		} else if i == 2 {
			return theme.Checkbox(&e.lightsCheckbox, "Lights").Layout(gtx)
		} else if i == 3 {
			return layout.Spacer{Width: 20}.Layout(gtx)
		} else if i == 4 {
			return theme.Checkbox(&e.bigCheckbox, "Big").Layout(gtx)
		} else if i == 5 {
			return layout.Spacer{Width: 20}.Layout(gtx)
		} else if i == 6 {
			return theme.Checkbox(&e.flipCheckbox, "Flip").Layout(gtx)
		} else if i == 7 {
			return layout.Spacer{Width: 60}.Layout(gtx)
		} else {
			return theme.Button(&e.cancelBtn, "Cancel").Layout(gtx)
		}
	})
	fmt.Println("dim", dim)
	return dim
}
