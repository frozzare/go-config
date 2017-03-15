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

		if err == ErrNoValueFound && len(def) > 0 {
			return value, nil
		}

		return value, err
	}

	return defaultValue(v, def...).(bool), nil
}

// MustBool returns a bool from the config file,
// it will panic if a error is created.
func MustBool(name string, def ...interface{}) bool {
	v, err := Bool(name, def...)

	if err != nil {
		panic(err)
	}

	return v
}

// Float returns a float64 from the config file.
func Float(name string, def ...interface{}) (float64, error) {
	v, err := config.value(name, floatType)

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
func MustFloat(name string, def ...interface{}) float64 {
	v, err := Float(name, def...)

	if err != nil {
		panic(err)
	}

	return v
}

// Int returns a int64 from the config file.
func Int(name string, def ...interface{}) (int64, error) {
	v, err := config.value(name, intType)

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
func MustInt(name string, def ...interface{}) int64 {
	v, err := Int(name, def...)

	if err != nil {
		panic(err)
	}

	return v
}

// List returns a slice of strings from the config file.
func List(name string, def ...interface{}) ([]string, error) {
	v, err := config.value(name, listType)

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
func MustList(name string, def ...interface{}) []string {
	v, err := List(name, def...)

	if err != nil {
		panic(err)
	}

	return v
}

// String returns a string from the config file.
func String(name string, def ...interface{}) (string, error) {
	v, err := config.value(name, stringType)

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
func MustString(name string, def ...interface{}) string {
	v, err := String(name, def...)

	if err != nil {
		panic(err)
	}

	return v
}

// Uint returns a unsigned int64 from the config file.
func Uint(name string, def ...interface{}) (uint64, error) {
	v, err := config.value(name, uintType)

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
func MustUint(name string, def ...interface{}) uint64 {
	v, err := Uint(name, def...)

	if err != nil {
		panic(err)
	}

	return v
}
