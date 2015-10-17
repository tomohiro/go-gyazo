Gyazo
================================================================================

[![Build Status](https://img.shields.io/travis/Tomohiro/go-gyazo.svg?style=flat-square)](https://travis-ci.org/Tomohiro/go-gyazo)
[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/Tomohiro/go-gyazo/gyazo)

Go library for accessing the Gyazo API


Usage
--------------------------------------------------------------------------------

### Create a client to accessing the Gyazo API

Import this package like this:

```go
import "github.com/Tomohiro/go-gyazo/gyazo"
```

Create a client with your Gyazo access token:

```go
client, err := gyazo.NewClient("your access token")
if err != nil {
  return err
}
```

### List

```go
list, _ := gyazo.List(&gyazo.ListOptions{Page: 1, PerPage: 50})
fmt.Println(list.Meta.TotalCount) // Total count of specified user's images
for _, img := range *list.Images {
    fmt.Println(img.PermalinkURL) // http://gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c
}
```

### Upload

```go
image, _ := gyazo.Upload("image.png")
fmt.Println(image.PermalinkURL) // http://gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c
```

### Delete

```go
result, _ := gyazo.Delete("8980c52421e452ac3355ca3e5cfe7a0c")
```
