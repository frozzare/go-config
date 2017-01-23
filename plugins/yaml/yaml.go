package yaml

import (
	"bytes"
	"io"

	yaml "gopkg.in/yaml.v2"

	config "github.com/frozzare/go-config"
)

func init() {
	config.RegisterFileType(".yml", Callback)
	config.RegisterFileType(".yaml", Callback)
}

func Callback(input io.Reader, output *map[string]interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	return yaml.Unmarshal(buf.Bytes(), &output)
}
