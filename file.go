package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
)

type fileType int

const (
	jsonType fileType = iota + 1
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
	}

	return nil
}
