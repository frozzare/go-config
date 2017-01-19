package config

import (
	"errors"
	"os"
	"strings"
)

type Env struct {
}

func (s *Env) Setup() error {
	return nil
}

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

	l := &Values{}
	l.Set(key, v)

	return l.Bool(key)
}

// Float returns a float64 or a error.
func (s *Env) Float(key string) (float64, error) {
	v, err := s.value(key)

	if err != nil {
		return 0.0, err
	}

	l := &Values{}
	l.Set(key, v)

	return l.Float(key)
}

// Int returns a int or a error.
func (s *Env) Int(key string) (int, error) {
	v, err := s.value(key)

	if err != nil {
		return 0, err
	}

	l := &Values{}
	l.Set(key, v)

	return l.Int(key)
}

// Get returns a interface or a error.
func (s *Env) Get(key string) (interface{}, error) {
	v, err := s.value(key)

	if err != nil {
		return nil, err
	}

	l := &Values{}
	l.Set(key, v)

	return l.Get(key)
}

// List returns a slice of strings or a error.
func (s *Env) List(key string) ([]string, error) {
	v, err := s.value(key)

	if err != nil {
		return []string{}, err
	}

	l := &Values{}
	l.Set(key, v)

	return l.List(key)
}

// String returns a string or a error.
func (s *Env) String(key string) (string, error) {
	v, err := s.value(key)

	if err != nil {
		return "", err
	}

	l := &Values{}
	l.Set(key, v)

	return l.String(key)
}

// Uint returns a unsigned int or a error.
func (s *Env) Uint(key string) (uint, error) {
	v, err := s.value(key)

	if err != nil {
		return 0, err
	}

	l := &Values{}
	l.Set(key, v)

	return l.Uint(key)
}
