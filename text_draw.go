package go_image_draw

import (
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

func NewTextDraw(fs []*truetype.Font, opts *truetype.Options) TextDraw {
	return newTextDraw(fs, opts)
}

type TextDraw interface {
	GetWidth(text string) float64
	GetTextWithWidth(text string, width float64) string
	WordWrap(text string, width float64) []string
	DrawString(im draw.Image, c color.Color, s string, x, y float64)
	DrawStringAnchored(im draw.Image, c color.Color, s string, x, y, ax, ay float64)
	DrawStringWrapped(im draw.Image, c color.Color, s string, x, y, ax, ay, width, lineSpacing float64, align Align)
}
