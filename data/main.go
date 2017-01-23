package main

import (
	"io"
	"log"
	"net/http"

	"github.com/frozzare/go-config"

	_ "github.com/frozzare/go-config/plugins/yaml"
)

// Example file

func main() {
	// Use file + env as a middlewares.
	config.Use(config.NewFromFile("config.yml"))
	config.Use(config.NewEnv())

	// Read and watch config.
	if err := config.ReadAndWatchFile("config.yml"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v, _ := config.String("name")
		io.WriteString(w, v)
	})

	http.ListenAndServe(":8899", nil)
}
