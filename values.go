package config

// Values middleware struct that handles predefined values.
type Values struct {
	values map[string]interface{}
}

// Setup returns a error if the middleware setup is failing.
func (s *Values) Setup() error {
	return nil
}

// NewFromValues creates a new values middleware.
func NewFromValues(values map[string]interface{}) *Values {
	return &Values{values}
}

// Bool returns a bool or a error.
func (s *Values) Bool(key string) (bool, error) {
	v, err := value(key, s.values)

	if err != nil {
		return false, err
	}

	return castBool(v)
}

// Float returns a float64 or a error.
func (s *Values) Float(key string) (float64, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0.0, err
	}

	return castFloat(v)
}

// Int returns a int or a error.
func (s *Values) Int(key string) (int, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0, err
	}

	return castInt(v)
}

// Get returns a interface or a error.
func (s *Values) Get(key string) (interface{}, error) {
	v, err := value(key, s.values)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// List returns a slice of strings or a error.
func (s *Values) List(key string) ([]string, error) {
	v, err := value(key, s.values)

	if err != nil {
		return []string{}, err
	}

	return castList(v)
}

// String returns a string or a error.
func (s *Values) String(key string) (string, error) {
	v, err := value(key, s.values)

	if err != nil {
		return "", err
	}

	return castString(v)
}

// Uint returns a unsigned int or a error.
func (s *Values) Uint(key string) (uint, error) {
	v, err := value(key, s.values)

	if err != nil {
		return 0, err
	}

	return castUint(v)
}

// ID returns the values struct identifier.
func (s *Values) ID() string {
	return "values"
}
