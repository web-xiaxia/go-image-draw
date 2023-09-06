package go_image_draw

import (
	"fmt"
	"testing"
)

func TestInitEmojiImageMap(t *testing.T) {
	image, b := GetEmojiImage("🌹")
	fmt.Println(image, b)
}
