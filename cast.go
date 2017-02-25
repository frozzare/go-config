package config

import "strings"

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
