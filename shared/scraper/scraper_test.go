package scraper

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/froostang/retail-therapy/shared/scraper/assets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAssetGetter is a mock implementation of the AssetGetter interface
type MockAssetGetter struct {
	mock.Mock
}

func (m *MockAssetGetter) GetImage(body []byte) (string, error) {
	args := m.Called(body)
	return args.String(0), args.Error(1)
}

func (m *MockAssetGetter) GetPrice(body []byte) (float64, error) {
	args := m.Called(body)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockAssetGetter) GetName(body []byte) (string, error) {
	args := m.Called(body)
	return args.String(0), args.Error(1)
}

func (m *MockAssetGetter) GetDescription(body []byte) (string, error) {
	args := m.Called(body)
	return args.String(0), args.Error(1)
}

// TestIsTargetURL tests the isTargetURL function
func TestIsTargetURL(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"http://www.target.com/some-product", true},
		{"https://www.target.com/another-product", true},
		{"http://www.example.com/some-product", false},
		{"http://target.com/some-product", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			actual := isTargetURL(tt.url)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// TestScrapeForImagePrice tests the ScrapeForImagePrice function
func TestScrapeForImagePrice(t *testing.T) {
	// Mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/expected-path" {
			t.Fatalf("expected to request '/expected-path', got: %s", r.URL.Path)
		}
		io.WriteString(w, `<html><body><img src="http://example.com/image.jpg"/><span class="price">current_retail":10.99</span></body></html>`)
	}))

	defer ts.Close()

	// Mock the AssetGetter interface
	mockGetter := new(MockAssetGetter)
	mockGetter.On("GetImage", mock.Anything).Return("http://example.com/image.jpg", nil)
	mockGetter.On("GetName", mock.Anything).Return("test", nil)
	mockGetter.On("GetPrice", mock.Anything).Return(10.99, nil)
	mockGetter.On("GetDescription", mock.Anything).Return("this is a test thingy", nil)

	tests := []struct {
		url              string
		expectedImageURL string
		expectedPrice    string
		expectedName     string
		expectedDesc     string
		expectedError    error
	}{
		{"http://www.target.com/expected-path", "http://example.com/image.jpg", "10.99", "test", "this is a test thingy", nil},
		{"http://www.not-target.com/expected-path", "", "", "", "", ErrBadURL},
		{ts.URL + "/unexpected-path", "", "", "", "", ErrBadURL},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			result, err := Scrape(tt.url, mockGetter)
			assert.Equal(t, tt.expectedImageURL, result.Image)
			assert.Equal(t, tt.expectedPrice, result.Price)
			assert.ErrorIs(t, tt.expectedError, err)
		})
	}
}

// TestGetPrice tests the GetPrice method of the Getter struct in the assets package
func TestGetPrice(t *testing.T) {
	tests := []struct {
		body          []byte
		expected      float64
		expectedError error
	}{
		{[]byte(`{\"current_retail\":10.99}`), 10.99, nil},
		{[]byte(`{\"current_retail\":0.99}`), 0.99, nil},
		{[]byte(`{\"current_retail\":"invalid"}`), 0, fmt.Errorf("price not found")},
		{[]byte(`{}`), 0, fmt.Errorf("price not found")},
	}

	for _, tt := range tests {
		t.Run(string(tt.body), func(t *testing.T) {
			getter := assets.Getter{}
			price, err := getter.GetPrice(tt.body)
			assert.Equal(t, tt.expected, price)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
