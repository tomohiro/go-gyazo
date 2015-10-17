package gyazo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/go-querystring/query"
)

// Image represents an uploaded image.
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

// ErrorResponse reports error caused by API request.
type ErrorResponse struct {
	Status  string
	Message string `json:"message"`
}

// Error returns the error response status and message.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %v", r.Status, r.Message)
}

// List represents the returned images and http headers from an API request.
type List struct {
	Meta   Meta
	Images *[]Image
}

// Meta represents the returned http headers from an API request.
type Meta struct {
	TotalCount  int
	CurrentPage int
	PerPage     int
	UserType    string
}

// ListOptions specifies the optional parameters to an API request.
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

// Upload uploads specified an image.
func (c *Client) Upload(path string) (*Image, error) {
	// Create multipart/form-data
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("imagedata", fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileBody)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Be aware that the URL is different from the other API.
	url := c.UploadEndpoint + "/api/upload"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, buildErrorResponse(res)
	}

	img := &Image{}
	if err = json.NewDecoder(res.Body).Decode(img); err != nil {
		return nil, err
	}

	return img, nil
}

// Delete deletes specified an image.
func (c *Client) Delete(id string) (*Image, error) {
	url := c.DefaultEndpoint + "/api/images/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, buildErrorResponse(res)
	}

	img := &Image{}
	if err = json.NewDecoder(res.Body).Decode(img); err != nil {
		return nil, err
	}

	return img, nil
}

// buildErrorResponse builds an error information from a HTTP response.
func buildErrorResponse(res *http.Response) error {
	er := &ErrorResponse{Status: res.Status}
	if err := json.NewDecoder(res.Body).Decode(er); err != nil {
		er.Message = err.Error()
	}
	return er
}

// createMeta creates a meta data from a HTTP response headers.
func createMeta(h http.Header) Meta {
	return Meta{
		TotalCount:  atoi(h["X-Total-Count"][0]),
		CurrentPage: atoi(h["X-Current-Page"][0]),
		PerPage:     atoi(h["X-Per-Page"][0]),
		UserType:    h["X-User-Type"][0],
	}
}
