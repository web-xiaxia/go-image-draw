package go_image_draw

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
)

//go:embed emoji.json
var emojiJson []byte
var ImageMap map[string]image.Image

func init() {
	_ = InitEmojiImageMap(true)
}

func InitEmojiImageMap(errContinue bool) error {
	var emojiImageBase64Map map[string]string
	if err := json.Unmarshal(emojiJson, &emojiImageBase64Map); err != nil {
		return errors.New("json unmarshal err")
	}

	ImageMap = make(map[string]image.Image, len(emojiImageBase64Map))
	for k, v := range emojiImageBase64Map {
		decodeString, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			if errContinue {
				continue
			}
			return errors.New(fmt.Sprintf("%s: base64 decode error, v：%s", k, v))
		}
		decode, _, err := image.Decode(bytes.NewBuffer(decodeString))
		if err != nil {
			if errContinue {
				continue
			}
			return errors.New(fmt.Sprintf("%s: image decode error, v：%s", k, v))
		}
		ImageMap[k] = decode
		//fmt.Println(k, len(k))
	}
	return nil
}
