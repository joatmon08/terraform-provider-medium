package testimage

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strings"
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

func GetBase64MD5(file []byte) string {
	md5sum := md5.Sum(file)
	md5 := strings.Replace(base64.StdEncoding.EncodeToString(md5sum[:16]), "=", "", -1)
	return md5
}
