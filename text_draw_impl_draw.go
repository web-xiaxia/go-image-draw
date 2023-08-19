package go_image_draw

import (
	splitter "github.com/SubLuLu/grapheme-splitter"
	"github.com/nfnt/resize"
	draw2 "golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

// DrawString draws the specified text at the specified point.
func (f *textDraw) DrawString(im draw.Image, c color.Color, s string, x, y float64) {
	f.DrawStringAnchored(im, c, s, x, y, 0, 0)
}

// DrawStringAnchored draws the specified text at the specified anchor point.
// The anchor point is x - w * ax, y - h * ay, where w, h is the size of the
// text. Use ax=0.5, ay=0.5 to center the text at the specified point.
func (f *textDraw) DrawStringAnchored(im draw.Image, c color.Color, s string, x, y, ax, ay float64) {
	w := f.GetWidth(s)
	x -= ax * w
	y += ay * f.getHeight()
	f.drawString(im, c, s, x, y)
}
func (f *textDraw) DrawStringWrapped(im draw.Image, c color.Color, s string, x, y, ax, ay, width, lineSpacing float64, align Align) {
	lines := f.WordWrap(s, width)

	// sync h formula with MeasureMultilineString
	h := float64(len(lines)) * f.fontHeight * lineSpacing
	h -= (lineSpacing - 1) * f.fontHeight

	x -= ax * width
	y -= ay * h
	switch align {
	case AlignLeft:
		ax = 0
	case AlignCenter:
		ax = 0.5
		x += width / 2
	case AlignRight:
		ax = 1
		x += width
	}
	ay = 1
	for _, line := range lines {
		f.DrawStringAnchored(im, c, line, x, y, ax, ay)
		y += f.fontHeight * lineSpacing
	}
}

func (f *textDraw) faceGlyphEmoji(dot fixed.Point26_6, emojiImage image.Image) (
	dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int26_6) {
	emojiResizeSize := uint(f.opts.Size * 0.8)
	emojiImageRet := resize.Resize(emojiResizeSize, emojiResizeSize, emojiImage, resize.NearestNeighbor)
	x := int(dot.X / 64)
	y := int(float64(dot.Y/64) - f.opts.Size*0.85)
	return image.Rectangle{
			Min: image.Pt(x, y),
			Max: image.Pt(x+int(f.opts.Size), y+int(f.opts.Size)),
		},
		emojiImageRet,
		image.Point{
			X: -int(f.opts.Size * 0.1),
			Y: -int(f.opts.Size * 0.1),
		},
		f.scale
}
func (f *textDraw) faceGlyph(dot fixed.Point26_6, ss rune) (
	dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int26_6, ok bool) {
	defer func() {
		if p := recover(); p != nil {
			//
		}
	}()
	for _, ff := range f.faceInfoList {
		dr, mask, maskp, advance, ok = ff.Face.Glyph(dot, ss)
		if !ok {
			continue
		}
		if dr.Dx() == 0 {
			continue
		}
		if ff.Font.Index(ss) == 0 {
			continue
		}
		return
	}

	return
}
func (f *textDraw) drawString(im draw.Image, c color.Color, s string, x, y float64) {
	fontSrc := image.NewUniform(c)
	dot := fixp(x, y)

	arr := splitter.Split(s)
	for _, ss := range arr {
		if emojiImage, ok := ImageMap[ss]; ok {
			dr, mask, maskp, advance := f.faceGlyphEmoji(dot, emojiImage)
			draw.Draw(im, dr, mask, maskp, draw.Over)

			// fmt.Printf("%s:-> %d\n", string(ss), advance)
			dot.X += advance
		} else {
			for _, sss := range ss {
				//fmt.Printf("%s: %d\n", string(ss), dot.X)
				dr, mask, maskp, advance, ok := f.faceGlyph(dot, sss)
				if !ok {
					continue
				}
				sr := dr.Sub(dr.Min)
				fx, fy := float64(dr.Min.X), float64(dr.Min.Y)
				s2d := f64.Aff3{1, 0, fx, 0, 1, fy}
				draw2.BiLinear.Transform(im, s2d, fontSrc, sr, draw2.Over, &draw2.Options{
					SrcMask:  mask,
					SrcMaskP: maskp,
				})
				// fmt.Printf("%s:-> %d\n", string(sss), advance)
				dot.X += advance
			}
		}
	}
}