package main

import (
	"io"
	"net/http"

	"github.com/frozzare/go-config"

	_ "github.com/frozzare/go-config/plugins/yaml"
)

// Example file

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
