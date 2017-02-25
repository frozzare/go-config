package config

import (
	"os"
	"strings"

	"github.com/goraz/cast"
)

// Env middleware struct that handles environment variables.
type Env struct {
}

// Setup returns a error if the middleware setup is failing.
func (s *Env) Setup() error {
	return nil
}

// NewEnv creates a new environment middleware.
func NewEnv() *Env {
	return &Env{}
}

func (s *Env) value(name string) string {
	name = strings.ToUpper(name)
	name = strings.Replace(name, ".", "_", -1)

	return os.Getenv(name)
}

// Bool returns a bool or a error.
func (s *Env) Bool(key string) (bool, error) {
	v := s.value(key)

	if v == "" {
		return false, nil
	}

	return cast.Bool(v)
}

// Float returns a float64 or a error.
func (s *Env) Float(key string) (float64, error) {
	v := s.value(key)

	if v == "" {
		return 0.0, nil
	}

	return cast.Float(v)
}

// Int returns a int or a error.
func (s *Env) Int(key string) (int64, error) {
	v := s.value(key)

	if v == "" {
		return 0, nil
	}

	return cast.Int(v)
}

// Get returns a interface or a error.
func (s *Env) Get(key string) (interface{}, error) {
	v := s.value(key)

	if v == "" {
		return nil, nil
	}

	return v, nil
}

// List returns a slice of strings or a error.
func (s *Env) List(key string) ([]string, error) {
	v := s.value(key)

	if v == "" {
		return []string{}, nil
	}

	return castList(v)
}

// String returns a string or a error.
func (s *Env) String(key string) (string, error) {
	v := s.value(key)

	if v == "" {
		return "", nil
	}

	return cast.String(v)
}

// Uint returns a unsigned int or a error.
func (s *Env) Uint(key string) (uint64, error) {
	v := s.value(key)

	if v == "" {
		return 0, nil
	}

	return cast.Uint(v)
}

// ID returns the values struct identifier.
func (s *Env) ID() string {
	return "env"
}
