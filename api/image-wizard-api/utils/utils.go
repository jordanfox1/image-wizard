package utils

import (
	"net/http"
	"strings"
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
