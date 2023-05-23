package theme

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
)

var fonts []text.FontFace = gofont.Collection()
var shaper *text.Shaper = text.NewShaper(fonts)
var defaultFont = text.Font{}

var yellow color.NRGBA = color.NRGBA{R: 248, G: 252, B: 232, A: 255}
var black color.NRGBA = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
