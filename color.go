package go_image_draw

import "image/color"

// NewColorWithHex sets the current color using a hex string. The leading pound
// sign (#) is optional. Both 3- and 6-digit variations are supported. 8 digits
// may be provided to set the alpha value as well.
func NewColorWithHex(x string) color.Color {
	r, g, b, a := parseHexColor(x)
	return NewColorWitRGBA255(r, g, b, a)
}

// NewColorWitRGBA255 sets the current color. r, g, b, a values should be between 0 and
// 255, inclusive.
func NewColorWitRGBA255(r, g, b, a int) color.Color {
	return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

// NewColorWitRGB255 sets the current color. r, g, b values should be between 0 and 255,
// inclusive. Alpha will be set to 255 (fully opaque).
func NewColorWitRGB255(r, g, b int) color.Color {
	return NewColorWitRGBA255(r, g, b, 255)
}

// NewColorWitRGBA sets the current color. r, g, b, a values should be between 0 and 1,
// inclusive.
func NewColorWitRGBA(r, g, b, a float64) color.Color {
	return color.NRGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: uint8(a * 255),
	}
}

// NewColorWitRGB sets the current color. r, g, b values should be between 0 and 1,
// inclusive. Alpha will be set to 1 (fully opaque).
func NewColorWitRGB(r, g, b float64) color.Color {
	return NewColorWitRGBA(r, g, b, 1)
}
