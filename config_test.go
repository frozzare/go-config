package config

import (
	"testing"

	assert "github.com/frozzare/go-assert"
)

func TestData(t *testing.T) {
	assert.Empty(t, config.Data())
	config.Set("hello", "world")
	assert.Equal(t, "world", config.Get("hello").(string))
}

func TestUse(t *testing.T) {
	Use(nil)
	assert.Equal(t, 0, len(config.Middlewares()))

	Use(NewEnv())
	assert.Equal(t, 1, len(config.Middlewares()))
}
