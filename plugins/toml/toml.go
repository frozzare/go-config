package toml

import (
	"io"

	"github.com/frozzare/go-config"
	"github.com/pelletier/go-toml"
)

func init() {
	config.RegisterFileType(".toml", Callback)
}

// Callback read input and decode it to output type or returns a error.
func Callback(input io.Reader, output *map[string]interface{}) error {
	tree, err := toml.LoadReader(input)

	if err != nil {
		return err
	}

	*output = tree.ToMap()

	return nil
}
