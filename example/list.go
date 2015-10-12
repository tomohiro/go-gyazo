package main

import (
	"fmt"
	"os"

	"github.com/Tomohiro/gyazo"
)

// GYAZO_TOKEN="Your Gyazo access token" go run list.go
func main() {
	token := os.Getenv("GYAZO_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "Environment variable `GYAZO_TOKEN` is empty.")
		os.Exit(1)
	}
	client, _ := gyazo.NewClient(token)
	images, _ := client.List(nil)
	for _, v := range *images {
		fmt.Println(v.PermalinkURL)
	}
}
