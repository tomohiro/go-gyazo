package gyazo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	version         = "0.1"
	defaultEndpoint = "https://api.gyazo.com"
	uploadEndpoint  = "https://upload.gyazo.com"
)

// Client manages communication with the Gyazo API
type Client struct {
	// Gyazo API access token.
	token string

	// client provides request to API endpoints.
	client *http.Client

	// DefaultEndpint is Gyazo API endpoint.
	DefaultEndpoint string

	// UploadEndpint is Gyazo upload API endpoint.
	UploadEndpoint string
}

// Image represents a uploaded image.
//
// Gyazo API docs: https://gyazo.com/api/docs/image
type Image struct {
	ID           string `json:"image_id"`
	PermalinkURL string `json:"permalink_url"`
	ThumbURL     string `json:"thumb_url"`
	URL          string `json:"url"`
	Type         string `json:"type"`
}

// NewClient returns a new Gyazo API client.
func NewClient(token string) (*Client, error) {
	var err error
	if token == "" {
		return nil, errors.New("access token is empty")
	}

	c := &Client{token, http.DefaultClient, defaultEndpoint, uploadEndpoint}
	return c, err
}

// List returns user images
func (c *Client) List() ([]Image, error) {
	var err error

	url := fmt.Sprintf("%s/api/images", c.DefaultEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	list := new([]Image)
	if err = json.NewDecoder(res.Body).Decode(&list); err != nil {
		return nil, err
	}
	return *list, nil
}
