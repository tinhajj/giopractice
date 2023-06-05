package component

import (
	"ui/theme"
	"ui/widget"

	"gioui.org/layout"
)

type Example struct {
	submitBtn widget.Clickable
	cancelBtn widget.Clickable

	on widget.Bool
}

func NewExample() Example {
	return Example{}
}

func (e *Example) Layout(gtx layout.Context) layout.Dimensions {
	list := layout.List{}

	return list.Layout(gtx, 5, func(gtx layout.Context, i int) layout.Dimensions {
		if i == 0 {
			return theme.Button(&e.submitBtn, "Submit").Layout(gtx)
		} else if i == 1 {
			return layout.Spacer{Width: 20}.Layout(gtx)
		} else if i == 2 {
			//return layout.Spacer{Width: 20}.Layout(gtx)
			return theme.Checkbox(&e.on).Layout(gtx)
		} else if i == 3 {
			return layout.Spacer{Width: 20}.Layout(gtx)
		} else {
			return theme.Button(&e.cancelBtn, "Cancel").Layout(gtx)
		}
	})
}
