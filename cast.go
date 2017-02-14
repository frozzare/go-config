package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func castBool(v interface{}) (bool, error) {
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

func castFloat(v interface{}) (float64, error) {
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

func castInt(v interface{}) (int, error) {
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

	if f, ok := v.(int64); ok {
		return int(f), nil
	}

	return 0, fmt.Errorf("unable to cast %T", v)
}

func castList(v interface{}) ([]string, error) {
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

func castString(v interface{}) (string, error) {
	f, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("unable to cast %T", v)
	}

	return f, nil
}

func castUint(v interface{}) (uint, error) {
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

	if f, ok := v.(int64); ok {
		return uint(f), nil
	}

	return 0, fmt.Errorf("unable to cast %T", v)
}
