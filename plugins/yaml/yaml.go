package yaml

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v2"

	"github.com/frozzare/go-config"
)

func init() {
	config.RegisterFileType(".yml", Callback)
	config.RegisterFileType(".yaml", Callback)
}

// Callback read input and decode it to output type or returns a error.
func Callback(input io.Reader, output *map[string]interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	return yaml.Unmarshal(buf.Bytes(), &output)
}
