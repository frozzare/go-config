package config

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func convert(input interface{}) interface{} {
	if _, ok := input.(map[interface{}]interface{}); !ok {
		return input
	}

	v2 := make(map[string]interface{})

	for key, value := range input.(map[interface{}]interface{}) {
		k := toString(key)
		v2[k] = value
	}

	return v2
}

func value(name string, values interface{}) (interface{}, error) {
	if values == nil {
		return nil, errors.New("Value does not exists")
	}

	parts := strings.Split(name, ".")

	if len(parts) > 1 {
		for i, part := range parts {
			name = part

			if i+1 >= len(parts) {
				continue
			}

			v, err := value(parts[i+1], convert(values).(map[string]interface{})[part])

			if err != nil {
				continue
			}

			if v != nil && reflect.ValueOf(v).Kind() != reflect.Map {
				return v, nil
			}

			values = v
		}
	}

	if _, ok := convert(values).(map[string]interface{}); !ok {
		return values, nil
	}

	val, ok := convert(values).(map[string]interface{})[name]

	if !ok {
		return nil, errors.New("Value does not exists")
	}

	return val, nil
}

func toString(value interface{}) string {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", value)
	}
}
