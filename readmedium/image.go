package readmedium

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Image struct {
	URL         string
	ContentType string
	MD5         string
}

func GetImage(imageURL string) (*Image, error) {
	var body []byte
	req, err := http.NewRequest("GET", imageURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	contentType := res.Header["Content-Type"][0]

	response, err := ioutil.ReadAll(res.Body)
	hash := fmt.Sprintf("%x", md5.Sum(response))

	image := &Image{
		URL:         imageURL,
		ContentType: contentType,
		MD5:         hash,
	}

	return image, nil
}
