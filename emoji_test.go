package go_image_draw

import (
	"testing"
)

func TestInitEmojiImageMap(t *testing.T) {
	err := initEmojiImageMap(false)
	if err != nil {
		t.Error(err)
	}
}
