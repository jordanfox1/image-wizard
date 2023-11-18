package routes_test

import (
	"bytes"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jordanfox1/image-wizard-api/api/image-wizard-api/routes"
)

func TestConvertEndpoint(t *testing.T) {

	// Create a new fiber app and setup routes
	app := fiber.New()
	routes.SetupRoutes(app)

	// Create sample images to test
	jpgImage := []byte("sample jpg image bytes")
	pngImage := []byte("sample png image bytes")

	// Test converting JPG to PNG
	req := httptest.NewRequest("POST", "/api/convert?format=png", bytes.NewReader(jpgImage))
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}

	pngImageBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !strings.Contains(string(pngImageBytes), "png") {
		t.Error("Expected png image in response but didn't get it")
	}

	// Test converting PNG to JPG
	req = httptest.NewRequest("POST", "/api/convert?format=jpg", bytes.NewReader(pngImage))
	resp, _ = app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200 but got %d", resp.StatusCode)
	}

	jpgImageBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !strings.Contains(string(jpgImageBytes), "jpg") {
		t.Error("Expected jpg image in response but didn't get it")
	}
}
