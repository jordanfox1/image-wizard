package handlers

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
)

func ConvertImage(image []byte, desiredFormat string) ([]byte, error) {
	if desiredFormat != "png" {
		return nil, fmt.Errorf("only png format supported")
	}

	// Decode JPEG image
	img, err := jpeg.Decode(bytes.NewReader(image))
	if err != nil {
		return nil, err
	}

	// Encode image as PNG
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
