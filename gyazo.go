package gyazo

import "errors"

const (
	version         = "0.1"
	defaultEndpoint = "https://api.gyazo.com"
	uploadEndpoint  = "https://upload.gyazo.com"
)

// A Client manages communication with the Gyazo API
type Client struct {
	// Gyazo API access token
	token string

	// DefaultEndpint is Gyazo API endpoint.
	DefaultEndpoint string

	// UploadEndpint is Gyazo upload API endpoint.
	UploadEndpoint string
}

// NewClient returns a new Gyazo API client.
func NewClient(token string) (*Client, error) {
	var err error
	if token == "" {
		return nil, errors.New("Required access token")
	}
	c := &Client{token, defaultEndpoint, uploadEndpoint}
	return c, err
}
