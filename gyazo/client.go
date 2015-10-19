package gyazo

import (
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	defaultEndpoint = "https://api.gyazo.com"
	uploadEndpoint  = "https://upload.gyazo.com"
)

// Client manages communication with the Gyazo API.
type Client struct {
	// client provides request to API endpoints.
	client *http.Client

	// DefaultEndpint is the Gyazo API endpoint.
	DefaultEndpoint string

	// UploadEndpint is the Gyazo upload API endpoint.
	UploadEndpoint string
}

// NewClient returns a new Gyazo API client.
func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("access token is empty")
	}

	// Create an OAuth2 client to authentication.
	oauthClient := oauth2.NewClient(
		oauth2.NoContext,
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
	)

	c := &Client{
		client:          oauthClient,
		DefaultEndpoint: defaultEndpoint,
		UploadEndpoint:  uploadEndpoint,
	}

	return c, nil
}
