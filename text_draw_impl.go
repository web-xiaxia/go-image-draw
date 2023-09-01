package go_image_draw

import (
	splitter "github.com/SubLuLu/grapheme-splitter"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"math"
	"strings"
)

func newTextDraw(fs []*truetype.Font, opts *truetype.Options) TextDraw {
	if opts.Size == 0 {
		opts.Size = 12
	}

	if opts.DPI == 0 {
		opts.DPI = 72
	}
	scale := fixed.Int26_6(0.5 + (opts.Size * opts.DPI * 64 / 72))

	sizeWith := math.Round(opts.Size*(opts.DPI/72)*100) / 100
	halfSizeWith := math.Round(sizeWith*100/2) / 100

	faceInfoList := make([]*FaceInfo, 0, len(fs))
	for _, f := range fs {
		xf := f
		faceInfoList = append(faceInfoList, &FaceInfo{
			Face: truetype.NewFace(xf, opts),
			Font: xf,
		})
	}
	fontFace := faceInfoList[0].Face
	fontHeight := float64(fontFace.Metrics().Height) / 64
	return &textDraw{
		firstFace:     fontFace,
		faceInfoList:  faceInfoList,
		opts:          opts,
		scale:         scale,
		sizeWithScale: float64(scale) * 0.8,
		sizeWith:      sizeWith,
		halfSizeWith:  halfSizeWith,
		fontHeight:    fontHeight,
	}
}

type FaceInfo struct {
	Face font.Face
	Font *truetype.Font
}

var _ TextDraw = (*textDraw)(nil)

type textDraw struct {
	firstFace     font.Face
	faceInfoList  []*FaceInfo
	opts          *truetype.Options
	scale         fixed.Int26_6
	sizeWithScale float64
	sizeWith      float64
	fontHeight    float64
	halfSizeWith  float64
}

func (f *textDraw) getWidthWithRune(r rune) float64 {
	advance, ok := f.firstFace.GlyphAdvance(r)
	if !ok {
		return 0
	}
	if float64(advance) >= f.sizeWithScale {
		return f.sizeWith
	}
	return f.halfSizeWith
}

func (f *textDraw) GetMetrics() font.Metrics {
	return f.firstFace.Metrics()
}

func (f *textDraw) GetWidth(text string) float64 {
	arr := splitter.Split(text)
	nowWidth := float64(0)
	for _, r := range arr {
		if _, ok := ImageMap[r]; ok {
			nowWidth += f.sizeWith
		} else {
			nowWidth += f.getWidth(r)
		}
	}
	return nowWidth
}
func (f *textDraw) GetHeight() float64 {
	return f.getHeight()
}
func (f *textDraw) getHeight() float64 {
	return f.fontHeight
}

func (f *textDraw) GetTextWithWidth(text string, width float64) string {
	if len(text) == 0 {
		return ""
	}
	arr := splitter.Split(text)
	nowWidth := float64(0)
	ret := make([]string, 0, len(arr))
	for _, r := range arr {
		if _, ok := ImageMap[r]; ok {
			nowWidth += f.sizeWith
			if nowWidth > width {
				break
			}
			ret = append(ret, r)
		} else {
			nowWidth += f.getWidth(r)
			if nowWidth > width {
				break
			}
			ret = append(ret, r)
		}
	}
	return strings.Join(ret, "")
}

// MeasureString returns the rendered width and height of the specified text
// given the current font face.
func (f *textDraw) getWidth(s string) float64 {
	for _, info := range f.faceInfoList {
		isThisFont := true
		for _, s2 := range s {
			if info.Font.Index(s2) == 0 {
				isThisFont = false
				break
			}
		}
		if isThisFont {
			d := &font.Drawer{
				Face: info.Face,
			}
			a := d.MeasureString(s)
			return float64(a >> 6)
		}
	}
	d := &font.Drawer{
		Face: f.firstFace,
	}
	a := d.MeasureString(s)
	return float64(a >> 6)
}

func (f *textDraw) MeasureMultilineString(s string, lineSpacing float64) (width, height float64) {
	lines := strings.Split(s, "\n")

	// sync h formula with DrawStringWrapped
	height = float64(len(lines)) * f.fontHeight * lineSpacing
	height -= (lineSpacing - 1) * f.fontHeight

	d := &font.Drawer{
		Face: f.firstFace,
	}

	// max width from lines
	for _, line := range lines {
		adv := d.MeasureString(line)
		currentWidth := float64(adv >> 6) // from gg.Context.MeasureString
		if currentWidth > width {
			width = currentWidth
		}
	}

	return width, height
}

func (f *textDraw) TextWrap(s string, width float64) []string {
	var result []string
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			result = append(result, line)
			continue
		}
		x := ""
		xWith := float64(0)
		fields := splitter.Split(line)
		for _, field := range fields {
			xWith += f.GetWidth(field)
			if xWith > width {
				if x == "" {
					result = append(result, field)
					x = ""
					xWith = 0
					continue
				} else {
					result = append(result, x)
					x = ""
					xWith = 0
				}
			}
			x += field
		}
		if x != "" {
			result = append(result, x)
		}
	}
	return result
}
