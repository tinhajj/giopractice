package theme

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
)

var fonts []text.FontFace = gofont.Collection()
var shaper *text.Shaper = text.NewShaper(fonts)
var defaultFont = text.Font{}

var Yellow color.NRGBA = color.NRGBA{R: 248, G: 252, B: 232, A: 255}
var Black color.NRGBA = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
var Background color.NRGBA = color.NRGBA{R: 105, G: 105, B: 105, A: 255}
