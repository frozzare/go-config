package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type fileType int

const (
	jsonType fileType = iota + 1
	yamlType
	naType
)

// File struct is a config file.
type File struct {
	gen    func() (io.Reader, error)
	typ    fileType
	values map[string]interface{}
}

// fileTypeFromPath returns right file type for the file.
func fileTypeFromPath(path string) fileType {
	if strings.HasSuffix(path, ".json") {
		return jsonType
	}

	if strings.HasSuffix(path, ".yml") || strings.HasPrefix(path, ".yaml") {
		return yamlType
	}

	return naType
}

// NewFromFile creates a new middleware from file.
func NewFromFile(path string) *Values {
	file := &File{
		gen: func() (io.Reader, error) {
			return os.Open(path)
		},
		typ: fileTypeFromPath(path),
	}

	err := file.Setup()
	if err != nil {
		return &Values{err: err}
	}

	return &Values{values: file.values}
}

// Setup returns a error if the middleware setup is failing.
func (s *File) Setup() error {
	if s.typ == naType {
		return errors.New("File type is not implemented")
	}

	r, err := s.gen()
	if err != nil {
		return err
	}

	s.values = make(map[string]interface{})

	switch s.typ {
	case jsonType:
		return json.NewDecoder(r).Decode(&s.values)
	case yamlType:
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		return yaml.Unmarshal(buf.Bytes(), &s.values)
	}

	return nil
}
