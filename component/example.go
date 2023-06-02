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
}

func (e *Example) Layout(gtx layout.Context) layout.Dimensions {
	th := material.NewTheme(gofont.Collection())
	return theme.Button(th, &e.submitBtn, "hello").Layout(gtx)
}
