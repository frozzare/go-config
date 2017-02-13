# Config [![Build Status](https://travis-ci.org/frozzare/go-config.svg?branch=master)](https://travis-ci.org/frozzare/go-config) [![GoDoc](https://godoc.org/github.com/frozzare/go-config?status.svg)](https://godoc.org/github.com/frozzare/go-config)

> WIP

Go package for dealing with configuration files, has built in support for environment variables and JSON files and support for YAML via plugin.

## Installation

```
$ go get github.com/frozzare/go-config
```
## Example

```go
package main

import (
	"io"
	"net/http"

	"github.com/frozzare/go-config"

	_ "github.com/frozzare/go-config/plugins/yaml"
)

func main() {
	// Use file + env as middlewares.
	config.Use(config.NewFromFile("config.yml"))
	config.Use(config.NewEnv())

	// Watch config (replaces first middleware since it's the same file path).
	config.WatchFile("config.yml")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v, _ := config.String("name")
		io.WriteString(w, v)
	})

	http.ListenAndServe(":8899", nil)
}

```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)