package component

import (
	"ui/theme"
	"ui/widget"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Example struct {
	submitBtn widget.Clickable
	cancelBtn widget.Clickable
}

func (e *Example) Layout(gtx layout.Context) layout.Dimensions {
	th := material.NewTheme(gofont.Collection())
	list := layout.List{}

	return list.Layout(gtx, 3, func(gtx layout.Context, i int) layout.Dimensions {
		if i == 0 {
			return theme.Button(th, &e.submitBtn, "Submit").Layout(gtx)
		} else if i == 1 {
			return layout.Spacer{Width: 20}.Layout(gtx)
		} else {
			return theme.Button(th, &e.cancelBtn, "Cancel").Layout(gtx)
		}
	})
}
