package go_image_draw

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image/color"
	"image/draw"
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

// NewTextDraw creates a new TextDraw.
// fs is a list of truetype.Fonts, Search for and use text within the font in order of drawing.
func NewTextDraw(fs []*truetype.Font, opts *truetype.Options) TextDraw {
	return newTextDraw(fs, opts)
}

type TextDraw interface {
	// GetWidth returns the width of the text.
	GetWidth(text string) float64
	// GetHeight returns the height of the text.
	GetHeight() float64
	// GetTextWithWidth returns the text that fits in the specified width.
	GetTextWithWidth(text string, width float64) string
	// WordWrap  returns the text that fits in the specified width.
	WordWrap(text string, width float64) []string
	// DrawString draws the specified text at the specified point.
	DrawString(im draw.Image, c color.Color, s string, x, y float64)
	// DrawStringAnchored draws the specified text at the specified anchor point.
	// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
	// text. Use ax=0.5, ay=0.5 to center the text at the specified point.
	DrawStringAnchored(im draw.Image, c color.Color, s string, x, y, ax, ay float64)
	// DrawStringWrapped word-wraps the specified string to the given max width
	// and then draws it at the specified anchor point using the given line
	// spacing and text alignment.
	DrawStringWrapped(im draw.Image, c color.Color, s string, x, y, ax, ay, width, lineSpacing float64, align Align) float64

	// DrawStringToDC draws the specified text at the specified point.
	DrawStringToDC(im *gg.Context, c color.Color, s string, x, y float64)
	// DrawStringAnchoredToDC draws the specified text at the specified anchor point.
	// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
	// text. Use ax=0.5, ay=0.5 to center the text at the specified point.
	DrawStringAnchoredToDC(im *gg.Context, c color.Color, s string, x, y, ax, ay float64)
	// DrawStringWrappedToDC word-wraps the specified string to the given max width
	// and then draws it at the specified anchor point using the given line
	// spacing and text alignment.
	DrawStringWrappedToDC(im *gg.Context, c color.Color, s string, x, y, ax, ay, width, lineSpacing float64, align Align) float64
}
