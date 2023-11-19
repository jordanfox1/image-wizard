package routes_test

import (
	"bytes"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/routes"
	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/utils"
)

func setupTestApp() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 20, // 10MB max request body size
	})
	routes.SetupRoutes(app)
	return app
}

var app = setupTestApp()

func getTestImage(imagePath string) []byte {
	image, err := os.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	return image
}

func TestConvertEndpoint(t *testing.T) {
	testCases := []struct {
		name                string
		inputFormat         string
		desiredFormat       string
		inputImagePath      string
		expectedContentType string
		expectedStatus      int
		expectedError       error
	}{
		{
			name:                "JPG to PNG",
			inputFormat:         "jpg",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/png",
			expectedStatus:      200,
			expectedError:       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputImage := getTestImage(tc.inputImagePath)
			req := httptest.NewRequest("POST", "/api/convert?format="+tc.desiredFormat, bytes.NewReader(inputImage))

			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status %d but got %d", tc.expectedStatus, resp.StatusCode)
			}

			outputImage, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()

			// Assert expected image format
			actualContentType := utils.GetContentType(outputImage)
			if tc.expectedContentType != actualContentType {
				t.Errorf("Expected type %s but got %s", tc.expectedContentType, actualContentType)
			}
		})
	}
}

func TestRootEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	body := string(bodyBytes)

	if body != "Hello, World 👋 - from image wizard api" {
		t.Errorf("Expected 'Hello, World 👋 - from image wizard api' but got '%s'", body)
	}
}
