package theme

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/unit"
)

// Palette contains the minimal set of colors that a widget may need to
// draw itself.
type Palette struct {
	// Bg is the background color atop which content is currently being
	// drawn.
	Bg color.NRGBA

	// Fg is a color suitable for drawing on top of Bg.
	Fg color.NRGBA

	// ContrastBg is a color used to draw attention to active,
	// important, interactive widgets such as buttons.
	ContrastBg color.NRGBA

	// ContrastFg is a color suitable for content drawn on top of
	// ContrastBg.
	ContrastFg color.NRGBA

	Yellow    color.NRGBA
	Black     color.NRGBA
	White     color.NRGBA
	LightGray color.NRGBA
}

var Theme struct {
	Shaper *text.Shaper
	Palette
	TextSize unit.Sp

	Font font.Font
}

// Initialize the theme
func init() {
	var fonts []text.FontFace = gofont.Collection()
	var shaper *text.Shaper = text.NewShaper(fonts)

	Theme = struct {
		Shaper *text.Shaper
		Palette
		TextSize unit.Sp
		Font     font.Font
	}{
		Shaper: shaper,
		Palette: Palette{
			Bg:        color.NRGBA{R: 105, G: 105, B: 105, A: 255},
			Yellow:    color.NRGBA{R: 248, G: 252, B: 232, A: 255},
			Black:     color.NRGBA{R: 0, G: 0, B: 0, A: 255},
			White:     color.NRGBA{R: 255, G: 255, B: 255, A: 255},
			LightGray: color.NRGBA{R: 211, G: 211, B: 211, A: 255},
		},
		TextSize: unit.Sp(16),
		Font:     font.Font{},
	}
}
