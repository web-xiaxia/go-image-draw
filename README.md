# go-image-draw

image draw multi font string text and emoji
![](https://raw.githubusercontent.com/web-xiaxia/go-image-draw/master/image.png)

## Installation

    go get -u github.com/web-xiaxia/go-image-draw

## Examples

```
dc := NewTextDraw([]*truetype.Font{font1, font2}, &truetype.Options{
    Size: 80,
})
dc.DrawString(img, color.White, "draw multi font text and emoji 😊", 10, 85)
```

## Reference code

[gg](https://github.com/fogleman/gg/): `gg` is a library for rendering 2D graphics in pure Go.
