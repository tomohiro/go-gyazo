package gyazo

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	version         = "0.0.1"
	defaultEndpoint = "https://api.gyazo.com"
	uploadEndpoint  = "https://upload.gyazo.com"
)

// Client manages communication with the Gyazo API.
type Client struct {
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
	Star         bool   `json:"star"`
	CreatedAt    string `json:"created_at"`
}

// List reporesents returned images and headers from `List` API.
type List struct {
	Meta   Meta
	Images *[]Image
}

// Meta represents returned http headers from a API request.
type Meta struct {
	TotalCount  int
	CurrentPage int
	PerPage     int
	UserType    string
}

// ListOptions specifies the optional parameters to `List` API.
type ListOptions struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// NewClient returns a new Gyazo API client.
func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("access token is empty")
	}

	// Create an OAuth2 client to authentication
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

// List lists the images the specified user.
func (c *Client) List(opts *ListOptions) (*List, error) {
	url := c.DefaultEndpoint + "/api/images"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Build and set query parameters
	if opts != nil {
		params, err := query.Values(opts)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = params.Encode()
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	list := &List{
		Images: new([]Image),
		Meta:   createMetaData(res.Header),
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	if err = json.NewDecoder(res.Body).Decode(&list.Images); err != nil {
		return nil, err
	}

	return list, nil
}

func createMetaData(h http.Header) Meta {
	return Meta{
		TotalCount:  atoi(h["X-Total-Count"][0]),
		CurrentPage: atoi(h["X-Current-Page"][0]),
		PerPage:     atoi(h["X-Per-Page"][0]),
		UserType:    h["X-User-Type"][0],
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
