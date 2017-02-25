package config

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

var (
	config Config

	// ErrNoValueFound it's the error when no value is found.
	ErrNoValueFound = errors.New("Value not found given key")
)

// Config is the struct for the config.
type Config struct {
	sync.RWMutex
	data        map[interface{}]interface{}
	middlewares []Middleware
	parsed      bool
}

// Get returns a config value by key or nil.
func (c *Config) Get(key interface{}) interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.data[key]
}

// Set will set a config value by key.
func (c *Config) Set(key interface{}, value interface{}) {
	c.Lock()
	defer c.Unlock()
	if c.data == nil {
		c.data = make(map[interface{}]interface{})
	}
	c.data[key] = value
}

// Data returns all config values.
func (c *Config) Data() map[interface{}]interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.data
}

// Middlewares returns a existing middlewares.
func (c *Config) Middlewares() []Middleware {
	c.RLock()
	defer c.RUnlock()
	return c.middlewares
}

// Reset will reset the config instance.
func Reset() {
	config = Config{}
}

// Use adds a middleware to the stack list.
func Use(middleware ...Middleware) {
	config.Lock()
	defer config.Unlock()

	// Remove unwanted middlewares.
	for i, m := range middleware {
		if m == nil {
			middleware = append(middleware[:i], middleware[i+1:]...)
		}
	}

	// Replace old middleware with a new one if the has the same id.
	for i1, m1 := range config.middlewares {
		for i2, m2 := range middleware {
			if m1.ID() == m2.ID() {
				config.middlewares[i1] = m2
				middleware = append(middleware[:i2], middleware[i2+1:]...)
			}
		}
	}

	config.middlewares = append(config.middlewares, middleware...)

	// Stop program if a middleware failes to setup.
	for _, middleware := range config.middlewares {
		err := middleware.Setup()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Forced to abort because %T is failing to setup: %v\n", middleware, err)
			os.Exit(1)
		}
	}

	// Be sure to reset config data since we adds a new middleware.
	config.data = make(map[interface{}]interface{})
}

// value loops through all middlewares to find the first value by key from a middleware.
func (c *Config) value(key string, typ valueType) (interface{}, error) {
	var value interface{}
	var err error

	if v := config.Get(key); v != nil {
		return v, nil
	}

	for _, middleware := range config.Middlewares() {
		switch typ {
		case boolType:
			o, err := middleware.Bool(key)

			if err != nil {
				continue
			}

			value = o
			break
		case floatType:
			o, err := middleware.Float(key)

			if err != nil {
				continue
			}

			value = o
			break
		case intType:
			o, err := middleware.Int(key)

			if err != nil {
				continue
			}

			value = o
			break
		case interfaceType:
			o, err := middleware.Get(key)

			if err != nil {
				continue
			}

			value = o
			break
		case listType:
			o, err := middleware.List(key)

			if err != nil {
				continue
			}

			value = o
			break
		case stringType:
			o, err := middleware.String(key)

			if err != nil {
				continue
			}

			value = o
			break
		case uintType:
			o, err := middleware.Uint(key)

			if err != nil {
				continue
			}

			value = o
			break
		}

		if value != nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, ErrNoValueFound
	}

	config.Set(key, value)

	return value, nil
}
