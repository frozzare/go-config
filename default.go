package config

import "github.com/goraz/cast"

func defaultValue(value interface{}, def ...interface{}) interface{} {
	switch value.(type) {
	case bool:
		v, err := cast.Bool(value)

		if err != nil && len(def) > 0 {
			return def[0]
		}

		if err != nil {
			return false
		}

		if v == false && len(def) > 0 {
			return def[0]
		}

		return v
	case float64:
		v, err := cast.Float(value)

		if (err != nil || v == 0.0) && len(def) > 0 {
			d, err := cast.Float(def[0])
			if err == nil {
				return d
			}
		}

		if err != nil {
			return 0.0
		}

		return v
	case int, int8, int16, int32, int64:
		v, err := cast.Int(value)
		if (err != nil || v == 0) && len(def) > 0 {
			d, err := cast.Int(def[0])

			if err == nil {
				return d
			}
		}

		if err != nil {
			return int64(0)
		}

		return int64(v)
	case string:
		v, err := cast.String(value)

		if err != nil && len(def) > 0 {
			return def[0]
		}

		if err != nil {
			return ""
		}

		if v == "" && len(def) > 0 {
			return def[0]
		}

		return v
	case []string:
		v, err := cast.StringSlice(value)

		if err != nil && len(def) > 0 {
			return def[0]
		}

		if err != nil {
			return []string{}
		}

		if len(v) == 0 && len(def) > 0 {
			return def[0]
		}

		return v
	case uint, uint8, uint16, uint32, uint64:
		v, err := cast.Uint(value)

		if (err != nil || v == 0) && len(def) > 0 {
			d, err := cast.Uint(def[0])

			if err == nil {
				return d
			}
		}

		if err != nil {
			return uint64(0)
		}

		return uint64(v)
	}

	return value
}
