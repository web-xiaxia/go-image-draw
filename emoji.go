package go_image_draw

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"io"
)

//go:embed emoji.json
var emojiJson []byte
var emojiImageBase64Map map[string]string

func init() {
	_ = initEmojiImageMap(true)
}

func initEmojiImageMap(errContinue bool) error {
	if len(emojiImageBase64Map) > 0 {
		return nil
	}
	emojiJsonReader, err := gzip.NewReader(bytes.NewBuffer(emojiJson))
	if err != nil {
		return err
	}
	emojiJson2, err := io.ReadAll(emojiJsonReader)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(emojiJson2, &emojiImageBase64Map); err != nil {
		return errors.New("json unmarshal err")
	}
	return nil
}
func IsEmoji(text string) bool {
	if len(text) == 0 {
		return false
	}
	if _, ok := emojiImageBase64Map[text]; ok {
		return true
	}
	return false
}

func GetEmojiImage(text string) (image.Image, bool) {
	if len(text) == 0 {
		return nil, false
	}
	if emojiImageBase64, ok := emojiImageBase64Map[text]; ok {
		if decodeString, err := base64.StdEncoding.DecodeString(emojiImageBase64); err == nil {
			if decode, _, err := image.Decode(bytes.NewBuffer(decodeString)); err == nil {
				return decode, true
			}
		}
	}
	return nil, false
}
