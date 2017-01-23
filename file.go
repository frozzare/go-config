package config

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"bytes"

	"github.com/fsnotify/fsnotify"
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
func NewFromFile(path string) Middleware {
	file := &File{
		gen: func() (io.Reader, error) {
			return os.Open(path)
		},
		callback: fileCallbackFromPath(path),
	}

	err := file.Setup()
	if err != nil {
		return &Values{err: err, id: path}
	}

	return &Values{values: file.values, id: path}
}

// NewFromBytes creates a new middleware from bytes as the given type, e.g: json.
func NewFromBytes(typ string, body []byte) Middleware {
	file := &File{
		gen: func() (io.Reader, error) {
			return bytes.NewReader(body), nil
		},
		callback: fileCallbackFromPath("." + typ),
	}

	err := file.Setup()
	if err != nil {
		return &Values{err: err, id: "bytes." + typ}
	}

	return &Values{values: file.values, id: "bytes." + typ}
}

// ReadAndWatchFile reads and watchs for changes in the configuration file.
func ReadAndWatchFile(filePath string) error {
	Use(NewFromFile(filePath))

	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		done := make(chan bool)
		go func() {
			for {
				select {
				case e := <-watcher.Events:
					if e.Op&fsnotify.Write == fsnotify.Write {
						Use(NewFromFile(filePath))
					}
				case <-watcher.Errors:
				}
			}
		}()

		err = watcher.Add(filePath)
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}()

	return nil
}

// RegisterFileType register a file type with a callback.
func RegisterFileType(ext string, callback fileCallback) {
	fileTypesMu.Lock()

	defer fileTypesMu.Unlock()

	if callback == nil {
		panic("config: RegisterFileType callback is nil")
	}

	if _, dup := fileTypes[ext]; dup {
		panic("config: RegisterFileType called twice for " + ext)
	}

	fileTypes[ext] = callback
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
