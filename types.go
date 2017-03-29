package config

import "fmt"

type valueType int

const (
	boolType valueType = iota + 1
	floatType
	intType
	interfaceType
	listType
	stringType
	uintType
)

// Bool returns a bool from the config file.
func Bool(key string, def ...interface{}) (bool, error) {
	v, err := config.value(key, boolType)

	if err != nil {
		value := defaultValue(false, def...).(bool)

		if err == ErrNoValueFound && len(def) > 0 {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(bool), nil
}

// MustBool returns a bool from the config file,
// it will panic if a error is created.
func MustBool(key string, def ...interface{}) bool {
	v, err := Bool(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}

// Float returns a float64 from the config file.
func Float(key string, def ...interface{}) (float64, error) {
	v, err := config.value(key, floatType)

	if err != nil {
		if err == ErrNoValueFound && len(def) > 0 {
			return defaultValue(0.0, def...).(float64), nil
		}

		return 0.0, err
	}

	return defaultValue(v, def...).(float64), nil
}

// MustFloat returns a float64 from the config file,
// it will panic if a error is created.
func MustFloat(key string, def ...interface{}) float64 {
	v, err := Float(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}

// Get returns the value for the given key from the config file as a interface.
func Get(key string, def ...interface{}) (interface{}, error) {
	v, err := config.value(key, interfaceType)

	if err != nil {
		if err == ErrNoValueFound && len(def) > 0 {
			return defaultValue(nil, def...), nil
		}

		return nil, err
	}

	return v, nil
}

// MustGet returns a interface from the config file,
// it will panic if a error is created.
func MustGet(key string, def ...interface{}) interface{} {
	v, err := Get(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}

// Int returns a int64 from the config file.
func Int(key string, def ...interface{}) (int64, error) {
	v, err := config.value(key, intType)

	if err != nil {
		if err == ErrNoValueFound && len(def) > 0 {
			return defaultValue(0, def...).(int64), nil
		}

		return 0, err
	}

	return defaultValue(v, def...).(int64), nil
}

// MustInt returns a int64 from the config file,
// it will panic if a error is created.
func MustInt(key string, def ...interface{}) int64 {
	v, err := Int(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}

// List returns a slice of strings from the config file.
func List(key string, def ...interface{}) ([]string, error) {
	v, err := config.value(key, listType)

	if err != nil {
		if err == ErrNoValueFound && len(def) > 0 {
			return defaultValue([]string{}, def...).([]string), nil
		}

		return []string{}, err
	}

	return defaultValue(v, def...).([]string), nil
}

// MustList returns a slice of strings from the config file,
// it will panic if a error is created.
func MustList(key string, def ...interface{}) []string {
	v, err := List(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}

// String returns a string from the config file.
func String(key string, def ...interface{}) (string, error) {
	v, err := config.value(key, stringType)

	if err != nil {
		if err == ErrNoValueFound && len(def) > 0 {
			return defaultValue("", def...).(string), nil
		}

		return "", err
	}

	return defaultValue(v, def...).(string), nil
}

// MustString returns a string from the config file,
// it will panic if a error is created.
func MustString(key string, def ...interface{}) string {
	v, err := String(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}

// Uint returns a unsigned int64 from the config file.
func Uint(key string, def ...interface{}) (uint64, error) {
	v, err := config.value(key, uintType)

	if err != nil {
		if err == ErrNoValueFound && len(def) > 0 {
			return defaultValue(uint64(0), def...).(uint64), nil
		}

		return uint64(0), err
	}

	return defaultValue(v, def...).(uint64), nil
}

// MustUint returns a unsigned int64 from the config file,
// it will panic if a error is created.
func MustUint(key string, def ...interface{}) uint64 {
	v, err := Uint(key, def...)

	if err != nil {
		panic(fmt.Errorf("%s: %s", err, key))
	}

	return v
}
