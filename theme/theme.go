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
	// Bg is the background color atop which content is currently being drawn.
	Bg color.NRGBA

	// Fg is a color suitable for drawing on top of Bg.
	Fg color.NRGBA

	Teal         color.NRGBA
	Teal100      color.NRGBA
	Yellow       color.NRGBA
	Black        color.NRGBA
	White        color.NRGBA
	LightGray200 color.NRGBA
	LightGray100 color.NRGBA
	Olive        color.NRGBA
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
			Bg:           color.NRGBA{R: 105, G: 105, B: 105, A: 255},
			Fg:           color.NRGBA{},
			Yellow:       color.NRGBA{R: 248, G: 252, B: 232, A: 255},
			Black:        color.NRGBA{R: 0, G: 0, B: 0, A: 255},
			White:        color.NRGBA{R: 255, G: 255, B: 255, A: 255},
			LightGray200: color.NRGBA{R: 210, G: 210, B: 210, A: 255},
			LightGray100: color.NRGBA{R: 235, G: 235, B: 235, A: 255},
			Olive:        color.NRGBA{R: 153, G: 153, B: 76, A: 255},
			Teal:         color.NRGBA{A: 255, R: 85, G: 170, B: 170},
			Teal100:      color.NRGBA{R: 234, G: 255, B: 255, A: 255},
		},
		TextSize: unit.Sp(16),
		Font:     font.Font{},
	}
}

// clamp1 limits v to range [0..1].
func clamp1(v float32) float32 {
	if v >= 1 {
		return 1
	} else if v <= 0 {
		return 0
	} else {
		return v
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
