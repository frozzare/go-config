package config

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

// Get returns the value for the given key from the config file as a interface.
func Get(name string) (interface{}, error) {
	v, err := config.value(name, interfaceType)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// Bool returns a bool from the config file.
func Bool(name string, def ...interface{}) (bool, error) {
	v, err := config.value(name, boolType)

	if err != nil {
		value := defaultValue(false, def...).(bool)

		if err == ErrNoValueFound && value != false {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(bool), nil
}

// Float returns a float64 from the config file.
func Float(name string, def ...interface{}) (float64, error) {
	v, err := config.value(name, floatType)

	if err != nil {
		value := defaultValue(0.0, def...).(float64)

		if err == ErrNoValueFound && value != 0.0 {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(float64), nil
}

// Int returns a integer from the config file.
func Int(name string, def ...interface{}) (int64, error) {
	v, err := config.value(name, intType)

	if err != nil {
		value := defaultValue(0, def...).(int64)

		if err == ErrNoValueFound && value != 0 {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(int64), nil
}

// List returns a slice of strings from the config file.
func List(name string, def ...interface{}) ([]string, error) {
	v, err := config.value(name, listType)

	if err != nil {
		value := defaultValue([]string{}, def...).([]string)

		if err == ErrNoValueFound && len(value) > 0 {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).([]string), nil
}

// String returns a string from the config file.
func String(name string, def ...interface{}) (string, error) {
	v, err := config.value(name, stringType)

	if err != nil {
		value := defaultValue("", def...).(string)

		if err == ErrNoValueFound && value != "" {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(string), nil
}

// Uint returns a unsigned int from the config file.
func Uint(name string, def ...interface{}) (uint64, error) {
	v, err := config.value(name, uintType)

	if err != nil {
		value := defaultValue(uint64(0), def...).(uint64)

		if err == ErrNoValueFound && value != uint64(0) {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(uint64), nil
}
