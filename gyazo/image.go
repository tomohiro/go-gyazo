package gyazo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// An Image represents a uploaded image.
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

// An ErrorResponse reports error caused by API request.
type ErrorResponse struct {
	Status  string
	Message string `json:"message"`
}

// Error returns error message.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %v", r.Status, r.Message)
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

	if res.StatusCode != http.StatusOK {
		return nil, buildErrorResponse(res)
	}

	list := &List{
		Images: new([]Image),
		Meta:   createMeta(res.Header),
	}

	if err = json.NewDecoder(res.Body).Decode(&list.Images); err != nil {
		return nil, err
	}

	return list, nil
}

// buildErrorResponse builds error information from responsed body.
func buildErrorResponse(res *http.Response) error {
	er := &ErrorResponse{Status: res.Status}
	if err := json.NewDecoder(res.Body).Decode(er); err != nil {
		er.Message = err.Error()
	}
	return er
}

// createMeta creates meta data from response headers.
func createMeta(h http.Header) Meta {
	return Meta{
		TotalCount:  atoi(h["X-Total-Count"][0]),
		CurrentPage: atoi(h["X-Current-Page"][0]),
		PerPage:     atoi(h["X-Per-Page"][0]),
		UserType:    h["X-User-Type"][0],
	}
}
