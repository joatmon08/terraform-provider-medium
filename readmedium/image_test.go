package readmedium

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func createTestImage() (*bytes.Buffer, error) {
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

func TestGetImage(t *testing.T) {
	buffer, err := createTestImage()
	if err != nil {
		t.Fatalf("unable to create image: %s", err)
	}
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/png")
			w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
			if _, err := w.Write(buffer.Bytes()); err != nil {
				t.Fatalf("unable to write image: %s", err)
			}
		}))
	defer ts.Close()

	imageURL := ts.URL
	expectedImage := &Image{
		URL:         imageURL,
		ContentType: "image/png",
		MD5:         fmt.Sprintf("%x", md5.Sum(buffer.Bytes())),
	}

	image, err := GetImage(imageURL)
	if err != nil {
		t.Fatalf("unable to access image: %s", err)
	}
	if image.URL != expectedImage.URL {
		t.Fatalf("image url is incorrect: %s", image.URL)
	}
	if image.MD5 != expectedImage.MD5 {
		t.Fatalf("image md5 is incorrect: %s", image.MD5)
	}
	if image.ContentType != expectedImage.ContentType {
		t.Fatalf("image content type is incorrect: %s", image.ContentType)
	}
}
