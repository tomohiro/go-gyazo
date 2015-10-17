package main

import (
	"fmt"
	"os"

	"github.com/Tomohiro/go-gyazo/gyazo"
)

// GYAZO_ACCESS_TOKEN="Your Gyazo access token" go run list.go
func main() {
	token := os.Getenv("GYAZO_ACCESS_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "Environment variable `GYAZO_ACCESS_TOKEN` is empty.")
		os.Exit(1)
	}

	client, _ := gyazo.NewClient(token)
	list, err := client.List(&gyazo.ListOptions{Page: 1, PerPage: 50})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	for _, v := range *list.Images {
		fmt.Printf("%+v \n", v)
	}
}
