package testimage

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
)

func Create() (*bytes.Buffer, error) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	var img image.Image = m
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		return nil, err
	}
	return buffer, nil
}