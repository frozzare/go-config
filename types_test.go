package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func init() {
	Use(NewFromFile("data/config.json"))
}

func TestTypeBool(t *testing.T) {
	v, err := Bool("bool")
	assert.Nil(t, err)
	assert.True(t, v)
}

func TestTypeFloat(t *testing.T) {
	v, err := Float("float")
	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestTypeInt(t *testing.T) {
	v, err := Int("int")
	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestTypeGet(t *testing.T) {
	v, err := Get("object")
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"name": "fredrik"}, v.(map[string]interface{}))
}

func TestTypeString(t *testing.T) {
	v, err := String("name")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestTypeList(t *testing.T) {
	v, err := List("names")
	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestTypeUint(t *testing.T) {
	v, err := Uint("uint")
	assert.Nil(t, err)
	assert.True(t, 1 == v)
}
