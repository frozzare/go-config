package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
)

// File struct is a config file.
type File struct {
	gen      func() (io.Reader, error)
	callback fileCallback
	values   map[string]interface{}
}

// fileCallback is used as a callback type for RegisterFileType.
type fileCallback func(input io.Reader, output *map[string]interface{}) error

var (
	fileTypesMu sync.RWMutex
	fileTypes   = make(map[string]fileCallback)
)

// RegisterFileType register a file type with a callback.
func RegisterFileType(ext string, callback fileCallback) {
	fileTypesMu.Lock()

	defer fileTypesMu.Unlock()

	if callback == nil {
		panic("store: RegisterFileType callback is nil")
	}

	if _, dup := fileTypes[ext]; dup {
		panic("config: RegisterFileType called twice for " + ext)
	}

	fileTypes[ext] = callback
}

// fileCallbackFromPath returns a callback for the file type or nil.
func fileCallbackFromPath(path string) fileCallback {
	if strings.HasSuffix(path, ".json") {
		return func(input io.Reader, output *map[string]interface{}) error {
			return json.NewDecoder(input).Decode(&output)
		}
	}

	for ext, callback := range fileTypes {
		if strings.HasSuffix(path, ext) {
			return callback
		}
	}

	return nil
}

// NewFromFile creates a new middleware from file.
func NewFromFile(path string) *Values {
	file := &File{
		gen: func() (io.Reader, error) {
			return os.Open(path)
		},
		callback: fileCallbackFromPath(path),
	}

	err := file.Setup()
	if err != nil {
		return &Values{err: err}
	}

	return &Values{values: file.values}
}

// Setup returns a error if the middleware setup is failing.
func (s *File) Setup() error {
	if s.callback == nil {
		return errors.New("File type is not implemented")
	}

	r, err := s.gen()
	if err != nil {
		return err
	}

	var values map[string]interface{}

	if err := s.callback(r, &values); err != nil {
		return err
	}

	s.values = values

	return nil
}
