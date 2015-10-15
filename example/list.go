package main

import (
	"fmt"
	"os"

	"github.com/Tomohiro/go-gyazo/gyazo"
)

// GYAZO_TOKEN="Your Gyazo access token" go run list.go
func main() {
	token := os.Getenv("GYAZO_TOKEN")
	if token == "" {
		fmt.Fprintf(os.Stderr, "Environment variable `GYAZO_TOKEN` is empty.")
		os.Exit(1)
	}
	client, _ := gyazo.NewClient(token)
	list, _ := client.List(&gyazo.ListOptions{Page: 1})
	for _, v := range *list.Images {
		fmt.Printf("%+v \n", v)
	}
}
