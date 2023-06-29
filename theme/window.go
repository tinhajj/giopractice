package theme

import (
	"image"
	"ui/widget"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type WindowStyle struct {
	Window *widget.Window
}

func Window(window *widget.Window) WindowStyle {
	return WindowStyle{
		Window: window,
	}
}

func (ws WindowStyle) Layout(gtx layout.Context) layout.Dimensions {
	return ws.Window.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return widget.OuterBorder{
			Border: widget.Border{
				Color:        Theme.Teal,
				CornerRadius: 0,
				Width:        unit.Dp(1),
			},
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			height := gtx.Dp(ws.Window.Height)
			width := gtx.Dp(ws.Window.Width)
			gtx.Constraints.Min = image.Point{}
			gtx.Constraints.Max = image.Point{X: width, Y: height}

			defer clip.Rect{Max: image.Point{X: width, Y: height}}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, Theme.Yellow)

			children := []layout.FlexChild{}

			title := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ws.Window.TitleBar(gtx, func(gtx layout.Context) layout.Dimensions {
					bg := widget.Background{Color: Theme.Teal100}
					return bg.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return Label(gtx, Theme.TextSize, ws.Window.Title)
						})
					})
				})
			})
			hr := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widget.HR{Width: unit.Dp(1), Color: Theme.Black}.Layout(gtx)
			})
			children = append(children, title, hr)

			if ws.Window.ContentWidget != nil {
				children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(3)).Layout(gtx, ws.Window.ContentWidget)
				}))
			}

			layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)

			return layout.Dimensions{
				Size: image.Point{X: width, Y: height},
			}
		})
	})
}
