package gyazo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, _ = NewClient("DUMMY_ACCESS_TOKEN")
	client.DefaultEndpoint = server.URL
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	c, err := NewClient("DUMMY_ACCESS_TOKEN")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	if actual, expected := c.DefaultEndpoint, defaultEndpoint; actual != expected {
		t.Errorf("NewClient DefaultEndpoint is %v, want %v", actual, expected)
	}

	if actual, expected := c.UploadEndpoint, uploadEndpoint; actual != expected {
		t.Errorf("NewClient UploadEndpoint is %v, want %v", actual, expected)
	}
}

func TestNewClient_EmptyAccessToken(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Error("Expected error to be returned")
	}
}
