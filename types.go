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
func Bool(name string) (bool, error) {
	v, err := config.value(name, boolType)

	if err != nil {
		return false, err
	}

	return v.(bool), nil
}

// Float returns a float64 from the config file.
func Float(name string) (float64, error) {
	v, err := config.value(name, floatType)

	if err != nil {
		return 0.0, err
	}

	return v.(float64), nil
}

// Int returns a integer from the config file.
func Int(name string) (int, error) {
	v, err := config.value(name, intType)

	if err != nil {
		return 0, err
	}

	return v.(int), nil
}

// List returns a slice of strings from the config file.
func List(name string) ([]string, error) {
	v, err := config.value(name, listType)

	if err != nil {
		return []string{}, err
	}

	return v.([]string), nil
}

// String returns a string from the config file.
func String(name string) (string, error) {
	v, err := config.value(name, stringType)

	if err != nil {
		return "", err
	}

	return v.(string), nil
}

// Uint returns a unsigned int from the config file.
func Uint(name string) (uint, error) {
	v, err := config.value(name, uintType)

	if err != nil {
		return 0, err
	}

	return v.(uint), nil
}
