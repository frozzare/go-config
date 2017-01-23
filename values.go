package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Values are a config middleware that can
// handle different type of values.
type Values struct {
	err    error
	values map[string]interface{}
}

// Setup returns a error if the middleware setup is failing.
func (s *Values) Setup() error {
	if s.err != nil {
		return s.err
	}

	if len(s.values) == 0 {
		return errors.New("Value does not exists")
	}

	return nil
}

// Set will add a value by key to the values map.
func (s *Values) Set(key string, value interface{}) {
	if s.values == nil {
		s.values = make(map[string]interface{})
	}

	s.values[key] = value
}

// Bool returns a bool from the values map or a error.
func (s *Values) Bool(key string) (bool, error) {
	v, err := value(key, s.values)
	if err != nil {
		return false, err
	}

	if x, ok := v.(string); ok {
		o, err := strconv.ParseBool(x)

		if err != nil {
			return false, err
		}

		return o, nil
	}

	b, ok := v.(bool)
	if !ok {
		return false, errors.New("unable to cast")
	}

	return b, nil
}

// Float returns a float64 from the values map or a error.
func (s *Values) Float(key string) (float64, error) {
	v, err := value(key, s.values)
	if err != nil {
		return 0, err
	}

	if x, ok := v.(string); ok {
		o, err := strconv.ParseFloat(x, 64)

		if err != nil {
			return 0.0, err
		}

		return o, nil
	}

	f, ok := v.(float64)
	if !ok {
		return 0.0, fmt.Errorf("unable to cast %T", v)
	}

	return f, nil
}

// Int returns a int from the values map or a error.
func (s *Values) Int(key string) (int, error) {
	v, err := value(key, s.values)
	if err != nil {
		return 0, err
	}

	if x, ok := v.(string); ok {
		o, err := strconv.Atoi(x)

		if err != nil {
			return 0, err
		}

		return o, nil
	}

	if f, ok := v.(float64); ok {
		return int(f), nil
	}

	if f, ok := v.(int); ok {
		return int(f), nil
	}

	return 0, fmt.Errorf("unable to cast %T", v)
}

// Get returns a interface from the values map or a error.
func (s *Values) Get(key string) (interface{}, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0, err
	}

	return v, nil
}

// List returns a slice of strings from the values map or a error.
func (s *Values) List(key string) ([]string, error) {
	v, err := value(key, s.values)
	if err != nil {
		return []string{}, err
	}

	if x, ok := v.(string); ok {
		l := strings.Split(x, ",")

		for i, p := range l {
			l[i] = strings.TrimSpace(p)
		}

		return l, nil
	}

	switch v.(type) {
	case []interface{}:
		v := v.([]interface{})
		result := make([]string, len(v))
		for i, item := range v {
			result[i] = toString(item)
		}
		return result, nil
	case []string:
		return v.([]string), nil
	}

	return []string{}, nil
}

// String returns a string from the values map or a error.
func (s *Values) String(key string) (string, error) {
	v, err := value(key, s.values)
	if err != nil {
		return "", err
	}

	f, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("unable to cast %T", v)
	}

	return f, nil
}

// Uint returns a unsigned int from the values map or a error.
func (s *Values) Uint(key string) (uint, error) {
	v, err := value(key, s.values)
	if err != nil {
		return 0, err
	}

	if x, ok := v.(string); ok {
		o, err := strconv.ParseUint(x, 10, 64)

		if err != nil {
			return 0, err
		}

		return uint(o), nil
	}

	if f, ok := v.(float64); ok {
		return uint(f), nil
	}

	if f, ok := v.(int); ok {
		return uint(f), nil
	}

	return 0, fmt.Errorf("unable to cast %T", v)
}
