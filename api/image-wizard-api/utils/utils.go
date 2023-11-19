package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

	"github.com/chai2010/webp"
)

func IsPNG(imageData []byte) bool {
	contentType := http.DetectContentType(imageData)

	return strings.HasSuffix(contentType, "png")
}

func IsJPG(imageData []byte) bool {
	contentType := http.DetectContentType(imageData)

	return strings.HasSuffix(contentType, "jpeg") || strings.HasSuffix(contentType, "jpg")
}

func GetContentType(imageData []byte) string {
	return http.DetectContentType(imageData)
}

func DecodeImage(inputImageData []byte, contentType string) (image.Image, error) {
	switch contentType {
	case "image/png":
		// Decode PNG image
		png, err := png.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return png, nil

	case "image/jpeg", "image/jpg":
		// Decode JPEG image
		jpg, err := jpeg.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return jpg, nil

	case "image/webp":
		// Decode WebP image
		webp, err := webp.Decode(bytes.NewReader(inputImageData))
		if err != nil {
			return nil, err
		}
		return webp, nil
	}

	return nil, fmt.Errorf("unsupported image format: %s", contentType)
}

func EncodeImage(desiredFormat string, img image.Image) ([]byte, error) {
	// Encode image to desired format
	switch desiredFormat {
	case "png":
		// Encode PNG image
		buf := new(bytes.Buffer)
		err := png.Encode(buf, img)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "jpeg", "jpg":
		// Encode JPEG image
		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case "webp":
		// Encode WebP image
		buf := new(bytes.Buffer)
		err := webp.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unsupported image format: %s", desiredFormat)
}
