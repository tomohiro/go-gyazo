Gyazo
================================================================================

[![Build Status](https://img.shields.io/travis/Tomohiro/go-gyazo.svg?style=flat-square)](https://travis-ci.org/Tomohiro/go-gyazo)
[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/Tomohiro/go-gyazo/gyazo)

Go library for accessing the Gyazo API


Usage
--------------------------------------------------------------------------------

### Upload

```go
gyazo, _ := gyazo.NewClient("your access token")
image, err := gyazo.Upload("image.png")
if err != nil {
  return err
}
fmt.Println(image.PermalinkURL) // https://gyazo.com/06643943e8dd0374f81bdeab02b71591
```
