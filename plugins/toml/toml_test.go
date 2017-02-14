package toml

import (
	"testing"

	assert "github.com/frozzare/go-assert"
	config "github.com/frozzare/go-config"
)

func init() {
	config.Use(config.NewFromFile("../../data/config.toml"))
}

func TestBool(t *testing.T) {
	v, err := config.Bool("bool")
	assert.Nil(t, err)
	assert.True(t, v)
}

func TestFloat(t *testing.T) {
	v, err := config.Float("float")
	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestInt(t *testing.T) {
	v, err := config.Int("int")
	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestGet(t *testing.T) {
	v, err := config.Get("object")
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"name": "fredrik"}, v)
}

func TestString(t *testing.T) {
	v, err := config.String("name")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestList(t *testing.T) {
	v, err := config.List("names")
	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestUint(t *testing.T) {
	v, err := config.Uint("uint")
	assert.Nil(t, err)
	assert.True(t, 1 == v)
}
