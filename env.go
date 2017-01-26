package config

import (
	"errors"
	"os"
	"strings"
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

func (s *Env) value(name string) (string, error) {
	name = strings.ToUpper(name)
	name = strings.Replace(name, ".", "_", -1)

	v := os.Getenv(name)
	if v == "" {
		return v, errors.New("Value does not exist")
	}

	return v, nil
}

// Bool returns a bool or a error.
func (s *Env) Bool(key string) (bool, error) {
	v, err := s.value(key)

	if err != nil {
		return false, err
	}

	return castBool(v)
}

// Float returns a float64 or a error.
func (s *Env) Float(key string) (float64, error) {
	v, err := s.value(key)

	if err != nil {
		return 0.0, err
	}

	return castFloat(v)
}

// Int returns a int or a error.
func (s *Env) Int(key string) (int, error) {
	v, err := s.value(key)

	if err != nil {
		return 0, err
	}

	return castInt(v)
}

// Get returns a interface or a error.
func (s *Env) Get(key string) (interface{}, error) {
	v, err := s.value(key)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// List returns a slice of strings or a error.
func (s *Env) List(key string) ([]string, error) {
	v, err := s.value(key)

	if err != nil {
		return []string{}, err
	}

	return castList(v)
}

// String returns a string or a error.
func (s *Env) String(key string) (string, error) {
	v, err := s.value(key)

	if err != nil {
		return "", err
	}

	return castString(v)
}

// Uint returns a unsigned int or a error.
func (s *Env) Uint(key string) (uint, error) {
	v, err := s.value(key)

	if err != nil {
		return 0, err
	}

	return castUint(v)
}

// ID returns the values struct identifier.
func (s *Env) ID() string {
	return "env"
}
