package theme

import (
	"gioui.org/font/gofont"
	"gioui.org/text"
)

var fonts []text.FontFace = gofont.Collection()
var shaper *text.Shaper = text.NewShaper(fonts)
var defaultFont = text.Font{}
