package routes_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
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
		expectError         bool
	}{
		// ---------- CONVERSIONS TO JPEG ----------
		{
			name:                "PNG to JPG",
			inputFormat:         "png",
			desiredFormat:       "jpg",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to JPG",
			inputFormat:         "webp",
			desiredFormat:       "jpg",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to JPG",
			inputFormat:         "webp",
			desiredFormat:       "jpg",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to JPG",
			inputFormat:         "webp",
			desiredFormat:       "jpg",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to JPG",
			inputFormat:         "webp",
			desiredFormat:       "jpg",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO PNG ----------
		{
			name:                "JPG to PNG",
			inputFormat:         "jpg",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to PNG",
			inputFormat:         "webp",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to PNG",
			inputFormat:         "webp",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to PNG",
			inputFormat:         "webp",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to PNG",
			inputFormat:         "webp",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO WEBP ----------
		{
			name:                "JPG to WEBP",
			inputFormat:         "jpg",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "PNG to WEBP",
			inputFormat:         "png",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to WEBP",
			inputFormat:         "jpg",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to WEBP",
			inputFormat:         "png",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to WEBP",
			inputFormat:         "png",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO BMP ----------
		{
			name:                "JPG to BMP",
			inputFormat:         "jpg",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "PNG to BMP",
			inputFormat:         "png",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to BMP",
			inputFormat:         "jpg",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to BMP",
			inputFormat:         "png",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to BMP",
			inputFormat:         "png",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO TIFF ----------
		{
			name:                "JPG to TIFF",
			inputFormat:         "jpg",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "PNG to TIFF",
			inputFormat:         "png",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to TIFF",
			inputFormat:         "jpg",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to TIFF",
			inputFormat:         "png",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to TIFF",
			inputFormat:         "png",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO GIF ----------
		{
			name:                "JPG to GIF",
			inputFormat:         "jpg",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "PNG to GIF",
			inputFormat:         "jpg",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to GIF",
			inputFormat:         "jpg",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to GIF",
			inputFormat:         "jpg",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to GIF",
			inputFormat:         "jpg",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- INVALID CONVERSIONS ----------
		{
			name:                "invalid desired format param",
			inputFormat:         "png",
			desiredFormat:       "1234",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusInternalServerError,
			expectError:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputImage := getTestImage(tc.inputImagePath)
			req := httptest.NewRequest("POST", fmt.Sprintf("/api/convert?format=%s", tc.desiredFormat), bytes.NewReader(inputImage))

			// Make request
			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatal(err)
			}

			// Assert status code and internal server error for expected errors
			if tc.expectError && tc.expectedStatus == http.StatusInternalServerError && resp.StatusCode == http.StatusInternalServerError {
				return
			}

			// Assert expected status code
			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status %d but got %d", tc.expectedStatus, resp.StatusCode)
			}

			// Assert expected image format
			outputImage, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}

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
