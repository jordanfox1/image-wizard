package handlers

import (
	"fmt"
	"strings"

	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/utils"
)

func ConvertImage(inputImageDataUrl string, desiredFormat string) ([]byte, error) {
	inputImageData, err := utils.GetImageDataFromDataURL(inputImageDataUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	inputImageContentType := utils.GetContentType(inputImageData)

	if strings.Contains(inputImageContentType, desiredFormat) {
		return nil, fmt.Errorf("input and output formats cannot be the same: %s", desiredFormat)
	}

	decodedImg, err := utils.DecodeImage(inputImageData, inputImageContentType)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	convertedImage, err := utils.EncodeImage(desiredFormat, decodedImg)
	if err != nil || convertedImage == nil {
		fmt.Println(err)
		return nil, err
	}

	return convertedImage, nil
}
