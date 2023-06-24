package component

import (
	"ui/theme"
	"ui/widget"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Example struct {
	List  widget.List
	Combo widget.Combo

	Editor widget.Editor

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

	return theme.List(&e.List).Layout(gtx, 11, func(gtx layout.Context, i int) layout.Dimensions {
		if i == 0 {
			return theme.Button(&e.submitBtn, "Submit").Layout(gtx)
		} else if i == 1 {
			return layout.Spacer{Height: 500, Width: 80}.Layout(gtx)
		} else if i == 2 {
			return theme.Checkbox(&e.lightsCheckbox, "Lights").Layout(gtx)
		} else if i == 3 {
			return theme.Combo(&e.Combo).Layout(gtx)
		} else if i == 4 {
			return theme.Checkbox(&e.bigCheckbox, "Big").Layout(gtx)
		} else if i == 5 {
			return layout.Spacer{Width: 40}.Layout(gtx)
		} else if i == 6 {
			return theme.Checkbox(&e.flipCheckbox, "Flip").Layout(gtx)
		} else if i == 7 {
			return layout.Spacer{Width: 60}.Layout(gtx)
		} else if i == 8 {
			return theme.Button(&e.cancelBtn, "Cancel").Layout(gtx)
		} else if i == 9 {
			return layout.Spacer{Width: 60}.Layout(gtx)
		} else {
			//return layout.Spacer{Height: 0, Width: 0}.Layout(gtx)
			//return layout.Dimensions{}
			return material.Editor(th, &e.Editor, "asd").Layout(gtx)
		}
	})
}
