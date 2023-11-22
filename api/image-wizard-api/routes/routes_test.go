package routes_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
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
			desiredFormat:       "jpeg",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to JPG",
			inputFormat:         "webp",
			desiredFormat:       "jpeg",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to JPG",
			inputFormat:         "bmp",
			desiredFormat:       "jpeg",
			inputImagePath:      "../test/images/bmp/sample.bmp",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to JPG",
			inputFormat:         "tiff",
			desiredFormat:       "jpeg",
			inputImagePath:      "../test/images/tiff/sample.tiff",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to JPG",
			inputFormat:         "gif",
			desiredFormat:       "jpeg",
			inputImagePath:      "../test/images/gif/sample.gif",
			expectedContentType: "image/jpeg",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO PNG ----------
		{
			name:                "JPG to PNG",
			inputFormat:         "jpeg",
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
			inputFormat:         "bmp",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/bmp/sample.bmp",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to PNG",
			inputFormat:         "tiff",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/tiff/sample.tiff",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to PNG",
			inputFormat:         "gif",
			desiredFormat:       "png",
			inputImagePath:      "../test/images/gif/sample.gif",
			expectedContentType: "image/png",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO WEBP ----------
		{
			name:                "JPG to WEBP",
			inputFormat:         "jpeg",
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
			inputFormat:         "bmp",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/bmp/sample.bmp",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to WEBP",
			inputFormat:         "tiff",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/tiff/sample.tiff",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to WEBP",
			inputFormat:         "gif",
			desiredFormat:       "webp",
			inputImagePath:      "../test/images/gif/sample.gif",
			expectedContentType: "image/webp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO BMP ----------
		{
			name:                "JPG to BMP",
			inputFormat:         "jpeg",
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
			inputFormat:         "webp",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to BMP",
			inputFormat:         "tiff",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/tiff/sample.tiff",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to BMP",
			inputFormat:         "gif",
			desiredFormat:       "bmp",
			inputImagePath:      "../test/images/gif/sample.gif",
			expectedContentType: "image/bmp",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},

		// ---------- CONVERSIONS TO TIFF ----------
		{
			name:                "JPG to TIFF",
			inputFormat:         "jpeg",
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
			inputFormat:         "webp",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to TIFF",
			inputFormat:         "bmp",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/bmp/sample.bmp",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "GIF to TIFF",
			inputFormat:         "gif",
			desiredFormat:       "tiff",
			inputImagePath:      "../test/images/gif/sample.gif",
			expectedContentType: "image/tiff",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		// ---------- CONVERSIONS TO GIF ----------
		{
			name:                "JPG to GIF",
			inputFormat:         "jpeg",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/jpg/foo.jpg",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "PNG to GIF",
			inputFormat:         "png",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/png/sample.png",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "WEBP to GIF",
			inputFormat:         "webp",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/webp/sample.webp",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "BMP to GIF",
			inputFormat:         "bmp",
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/bmp/sample.bmp",
			expectedContentType: "image/gif",
			expectedStatus:      http.StatusOK,
			expectError:         false,
		},
		{
			name:                "TIFF to GIF",
			inputFormat:         "tiff", // Corrected from "jpg" to "tiff"
			desiredFormat:       "gif",
			inputImagePath:      "../test/images/tiff/sample.tiff",
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
			var imageData []byte = getTestImage(tc.inputImagePath)
			base64ImageData := base64.StdEncoding.EncodeToString(imageData)
			inputFileName := filepath.Base(tc.inputImagePath)

			// Create form values
			formValues := map[string]string{
				"fileName": inputFileName,
				"image":    "data:image/" + tc.inputFormat + ";base64," + base64ImageData,
			}

			// Create a buffer to write the request body
			var body bytes.Buffer
			writer := multipart.NewWriter(&body)

			// Add form values to the request body
			for key, value := range formValues {
				writer.WriteField(key, value)
			}

			// Close the multipart writer
			writer.Close()

			// Create the request with the form-encoded body and image data
			req := httptest.NewRequest("POST", "/api/convert?format="+tc.desiredFormat, &body)
			req.Header.Set("Content-Type", writer.FormDataContentType())

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

			var responseMap map[string]interface{}
			if err := json.Unmarshal(outputImage, &responseMap); err != nil {
				t.Fatal(err)
			}

			// Check the fields in the response
			if status, ok := responseMap["status"].(float64); ok && int(status) == http.StatusOK {
				if dataURL, ok := responseMap["dataURL"].(string); ok {
					// Check that dataURL contains an image of the expected format
					imageData, _ := utils.GetImageDataFromDataURL(dataURL)
					actualType := utils.GetContentType(imageData)

					if actualType != "image/"+tc.desiredFormat {
						t.Errorf("Expected image format %s but dataURL contains image of type %s", tc.desiredFormat, actualType)
					}
				}

				if fileName, ok := responseMap["fileName"].(string); ok {
					expectedFileNamePrefix := getCharsBeforeDot(inputFileName)
					expectedFileName := expectedFileNamePrefix + "." + tc.desiredFormat
					if fileName != expectedFileName {
						t.Errorf("Expected filename %s but got %s as the filename", expectedFileName, fileName)
					}
				}
			} else {
				t.Errorf("Expected status %d but got %f", http.StatusOK, status)
			}
		})
	}
}

func getCharsBeforeDot(s string) string {
	i := strings.LastIndex(s, ".")
	if i == -1 {
		return s
	}
	return s[:i]
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
