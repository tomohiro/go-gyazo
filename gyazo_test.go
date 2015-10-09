package gyazo

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("DUMMY_ACCESS_TOKEN")

	if got, want := c.DefaultEndpoint, defaultEndpoint; got != want {
		t.Errorf("NewClient DefaultEndpoint is %v, want %v", got, want)
	}

	if got, want := c.UploadEndpoint, uploadEndpoint; got != want {
		t.Errorf("NewClient UploadEndpoint is %v, want %v", got, want)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestNewClientWithEmptyAccessToken(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Error("Expected error to be returned.")
	}
}
