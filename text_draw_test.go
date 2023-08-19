package go_image_draw

import (
	_ "embed"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestDrawEmoji(t *testing.T) {
	//emoji.InitEmojiImageMap(false)
	img := image.NewRGBA(image.Rectangle{
		Min: image.Pt(0, 0),
		Max: image.Pt(1200, 110),
	})

	regularFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		t.Fatal(err)
	}

	dc := NewTextDraw([]*truetype.Font{regularFont}, &truetype.Options{
		Size: 80,
	})
	dc.DrawString(img, color.White, "draw multi font text and emoji 😊", 10, 85)

	file, err := os.Create("testDrawEmoji.png")
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDrawStringWrapped(t *testing.T) {
	//emoji.InitEmojiImageMap(false)
	img := image.NewRGBA(image.Rectangle{
		Min: image.Pt(0, 0),
		Max: image.Pt(200, 400),
	})

	regularFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		t.Fatal(err)
	}

	dc := NewTextDraw([]*truetype.Font{regularFont}, &truetype.Options{
		Size: 20,
	})
	dc.DrawStringWrapped(img, color.White, "12345678901234567890", 0, 70, 0, 0, 190, 1, AlignCenter)

	file, err := os.Create("testDrawStringWrapped.png")
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		t.Fatal(err)
	}
}
