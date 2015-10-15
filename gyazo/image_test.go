package gyazo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/images", func(w http.ResponseWriter, r *http.Request) {
		// Set response headers
		// This example is response header that from https://gyazo.com/api/docs/image.
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total-Count", "350")
		w.Header().Set("X-Current-Page", "1")
		w.Header().Set("X-Per-Page", "20")
		w.Header().Set("X-User-Type", "lite")

		// Set 200 OK as HTTP status code.
		w.WriteHeader(http.StatusOK)

		// Set response body
		// This example is response body that from https://gyazo.com/api/docs/image.
		fmt.Fprintln(w, `[
			{
				"image_id": "8980c52421e452ac3355ca3e5cfe7a0c",
				"permalink_url": "http://gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c",
				"thumb_url": "https://i.gyazo.com/thumb/afaiefnaf.png",
				"url": "https://i.gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c.png",
				"type": "png",
				"star": true,
				"created_at": "2014-05-21 14:23:10+0900"
			}
		]`)
	})

	res, err := client.List(nil)
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}

	actual := res.Images
	expected := &[]Image{{
		ID:           "8980c52421e452ac3355ca3e5cfe7a0c",
		PermalinkURL: "http://gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c",
		ThumbURL:     "https://i.gyazo.com/thumb/afaiefnaf.png",
		URL:          "https://i.gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c.png",
		Type:         "png",
		Star:         true,
		CreatedAt:    "2014-05-21 14:23:10+0900",
	}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("List returned %+v, want %+v", actual, expected)
	}
}

func TestList_InvalidToken(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/images", func(w http.ResponseWriter, r *http.Request) {
		// Set 401 Unauthorized as HTTP status code.
		w.WriteHeader(http.StatusUnauthorized)

		// Set response body
		fmt.Fprintln(w, `{"message": "You are not authorized"}`)
	})

	_, err := client.List(nil)
	if err == nil {
		t.Error("Expected error to be returned")
	}

	expected := "401 Unauthorized: You are not authorized"
	actual := err.Error()
	if actual != expected {
		t.Errorf("ErrorResponse message is %v, want %v", actual, expected)
	}
}
