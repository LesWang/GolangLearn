package main

import(
    "image"
    "image/color"
)

type Image struct {

}

func (myImage Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (myImage Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 1, 1)
}

func (myImage Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}
