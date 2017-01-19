package config

import (
	"fmt"
	"os"
	"sync"
)

var config Config

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

// Use adds a middleware to the stack list.
func Use(middleware ...Middleware) {
	config.Lock()
	defer config.Unlock()
	config.middlewares = append(config.middlewares, middleware...)
}

func init() {
	for _, middleware := range config.Middlewares() {
		err := middleware.Setup()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Forced to abort because %T is failing to setup: %v\n", middleware, err)
			os.Exit(1)
		}
	}
}

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
	}

	if err != nil {
		return nil, err
	}

	config.Set(key, value)

	return value, nil
}
