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

	fmt.Println("EXAMPLE LAYOUT")
	fmt.Println("LAYOUT CONSTRAINT", gtx.Constraints)

	return material.List(th, &e.list).Layout(gtx, 9, func(gtx layout.Context, i int) layout.Dimensions {
		if i == 0 {
			dim := theme.Button(&e.submitBtn, "Submit").Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 1 {
			dim := layout.Spacer{Width: 80}.Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 2 {
			dim := theme.Checkbox(&e.lightsCheckbox, "Lights").Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 3 {
			dim := layout.Spacer{Width: 20}.Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 4 {
			dim := theme.Checkbox(&e.bigCheckbox, "Big").Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 5 {
			dim := layout.Spacer{Width: 40}.Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 6 {
			dim := theme.Checkbox(&e.flipCheckbox, "Flip").Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else if i == 7 {
			dim := layout.Spacer{Width: 60}.Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		} else {
			dim := theme.Button(&e.cancelBtn, "Cancel").Layout(gtx)
			fmt.Printf("%d constraint: %+v dim: %+v\n", i, gtx.Constraints, dim)
			return dim
		}
	})
}
