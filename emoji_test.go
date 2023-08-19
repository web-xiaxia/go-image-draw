package go_image_draw

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInitEmojiImageMap(t *testing.T) {
	err := InitEmojiImageMap(false)
	if err != nil {
		t.Error(err)
	}
}

func TestEmojiJson(t *testing.T) {
	dirName := "/Users/xiaxia/web/dev/w/xemoji/image"
	dirEntries, err := os.ReadDir(dirName)
	if err != nil {
		return
	}
	ret := make(map[string]string, len(dirEntries))
	for _, dirEntry := range dirEntries {
		nameArr := strings.Split(dirEntry.Name(), ".")
		if len(nameArr) != 2 || nameArr[1] != "png" {
			continue
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", dirName, dirEntry.Name()))
		if err != nil {
			return
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(file)
		all, err := ioutil.ReadAll(file)
		if err != nil {
			return
		}
		reader, err := vips.NewImageFromReader(bytes.NewBuffer(all))
		if err != nil {
			panic(fmt.Sprintf("%s:%+v", nameArr[0], err))
		}
		pngImage, _, err := reader.ExportPng(&vips.PngExportParams{
			StripMetadata: false,
			Quality:       100,
		})
		if err != nil {
			panic(fmt.Sprintf("%s:%+v", nameArr[0], err))
		}
		ret[nameArr[0]] = base64.StdEncoding.EncodeToString(pngImage)
	}
	marshal, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("emoji.json", marshal, 666)
	if err != nil {
		panic(err)
	}

}
