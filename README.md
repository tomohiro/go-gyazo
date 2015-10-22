go-gyazo
================================================================================

[![Build Status](https://img.shields.io/travis/Tomohiro/go-gyazo.svg?style=flat-square)](https://travis-ci.org/Tomohiro/go-gyazo)
[![Coverage Status](https://img.shields.io/coveralls/Tomohiro/go-gyazo.svg?style=flat-square)](https://coveralls.io/github/Tomohiro/go-gyazo)
[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/Tomohiro/go-gyazo/gyazo)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/Tomohiro/go-gyazo/blob/master/LICENSE)

go-gyazo is a library for Go for accessing the [Gyazo API](https://gyazo.com/api).


Usage
--------------------------------------------------------------------------------

### Create a client to accessing the Gyazo API

Import this package like this:

```go
import "github.com/Tomohiro/go-gyazo/gyazo"
```

Create a client with your [Gyazo access token](https://gyazo.com/oauth/applications):

```go
gyazo, err := gyazo.NewClient("your access token")
if err != nil {
	panic(err)
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
file, _ := os.Open("/your/image/file.png")
defer file.Close()
image, _ := gyazo.Upload(file)
fmt.Println(image.PermalinkURL) // http://gyazo.com/8980c52421e452ac3355ca3e5cfe7a0c
```

### Delete

```go
result, _ := gyazo.Delete("8980c52421e452ac3355ca3e5cfe7a0c")
```

For complete usage of go-gyazo, see the full [package docs](https://godoc.org/github.com/Tomohiro/go-gyazo/gyazo).


Contributing
--------------------------------------------------------------------------------

Please check out the [CONTIRBUTING](CONTRIBUTING.md) guideline.


LICENSE
--------------------------------------------------------------------------------

&copy; 2015 Tomohiro TAIRA.

This project is licensed under the MIT license. See [LICENSE](LICENSE) for details.
