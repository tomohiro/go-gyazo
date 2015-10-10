package gyazo

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("DUMMY_ACCESS_TOKEN")
	if err != nil {
		t.Error(err)
	}

	if got, want := c.DefaultEndpoint, defaultEndpoint; got != want {
		t.Errorf("NewClient DefaultEndpoint is %v, want %v", got, want)
	}

	if got, want := c.UploadEndpoint, uploadEndpoint; got != want {
		t.Errorf("NewClient UploadEndpoint is %v, want %v", got, want)
	}
}

func TestNewClientWithEmptyAccessToken(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Error("Expected error to be returned.")
	}
}

func TestList(t *testing.T) {
	c, _ := NewClient("DUMMY_ACCESS_TOKEN")
	_, err := c.List()
	if err != nil {
		t.Error(err)
	}
}
