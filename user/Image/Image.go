package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{

}

func (imag *Image) ColorModel()  color.Model{
	return color.RGBAModel
}

func (imag *Image) Bounds() image.Rectangle{
	return 	image.Rect(0, 0, 100, 100)
}

func (imag *Image) At(x, y int)  color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}



func main() {
	m := &Image{}
	pic.ShowImage(m)
}
