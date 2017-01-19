package config

// Middleware is the interface that external middlewares must implement.
type Middleware interface {
	Bool(name string) (bool, error)
	Float(name string) (float64, error)
	Int(name string) (int, error)
	Get(name string) (interface{}, error)
	List(name string) ([]string, error)
	Setup() error
	String(name string) (string, error)
	Uint(name string) (uint, error)
}
