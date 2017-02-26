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
	"github.com/goraz/cast"
)

// File struct is a config file.
type File struct {
	callback fileCallback
	gen      func() (io.Reader, error)
	path     string
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

// NewFromFile creates a new middleware from file. Optional bool argument to watch file.
func NewFromFile(path string, watch ...bool) Middleware {
	file := &File{
		callback: fileCallbackFromPath(path),
		gen: func() (io.Reader, error) {
			return os.Open(path)
		},
		path: path,
	}

	err := file.Setup()
	if err != nil {
		return nil
	}

	if len(watch) > 0 && watch[0] {
		WatchFile(path)
	}

	return file
}

// NewFromBytes creates a new middleware from bytes as the given type, e.g: json.
func NewFromBytes(typ string, body []byte) Middleware {
	file := &File{
		callback: fileCallbackFromPath("." + typ),
		gen: func() (io.Reader, error) {
			return bytes.NewReader(body), nil
		},
		path: "bytes." + typ,
	}

	err := file.Setup()
	if err != nil {
		return nil
	}

	return file
}

// ReadAndWatchFile reads and watches file for changes and reload the configuration file.
func ReadAndWatchFile(path string) {
	Use(NewFromFile(path))
	WatchFile(path)
}

// WatchFile watches file for changes and reload the configuration file.
func WatchFile(path string) {
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
						Use(NewFromFile(path))
					}
				case <-watcher.Errors:
				}
			}
		}()

		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}()
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

// Bool returns a bool or a error.
func (s *File) Bool(key string) (bool, error) {
	v, err := value(key, s.values)

	if err != nil {
		return false, err
	}

	return cast.Bool(v)
}

// Float returns a float64 or a error.
func (s *File) Float(key string) (float64, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0.0, err
	}

	return cast.Float(v)
}

// Int returns a int or a error.
func (s *File) Int(key string) (int64, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0, err
	}

	return cast.Int(v)
}

// Get returns a interface or a error.
func (s *File) Get(key string) (interface{}, error) {
	v, err := value(key, s.values)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// List returns a slice of strings or a error.
func (s *File) List(key string) ([]string, error) {
	v, err := value(key, s.values)

	if err != nil {
		return []string{}, err
	}

	return castList(v)
}

// String returns a string or a error.
func (s *File) String(key string) (string, error) {
	v, err := value(key, s.values)

	if err != nil {
		return "", err
	}

	return cast.String(v)
}

// Uint returns a unsigned int or a error.
func (s *File) Uint(key string) (uint64, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0, err
	}

	return cast.Uint(v)
}

// ID returns the values struct identifier.
func (s *File) ID() string {
	return s.path
}
