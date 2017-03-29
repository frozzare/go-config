package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestTypeBool(t *testing.T) {
	Reset()

	v, err := Bool("bool")
	assert.NotNil(t, err)
	assert.False(t, v)

	config.Set("bool", true)

	v, err = Bool("bool")
	assert.Nil(t, err)
	assert.True(t, v)
}

func TestTypeMustBool(t *testing.T) {
	Reset()

	v := MustBool("bool", false)
	assert.False(t, v)

	config.Set("bool", true)

	v = MustBool("bool")
	assert.True(t, v)
}

func TestTypeFloat(t *testing.T) {
	Reset()

	v, err := Float("float")
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, v)

	config.Set("float", 12.13)

	v, err = Float("float")
	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestTypeMustFloat(t *testing.T) {
	Reset()

	v := MustFloat("float", 12.00)
	assert.Equal(t, 12.00, v)

	config.Set("float", 12.13)

	v = MustFloat("float")
	assert.Equal(t, 12.13, v)
}

func TestTypeGet(t *testing.T) {
	Reset()

	v, err := Get("object")
	assert.NotNil(t, err)
	assert.Nil(t, v)

	config.Set("object", map[string]interface{}{"name": "fredrik"})

	v, err = Get("object")

	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"name": "fredrik"}, v.(map[string]interface{}))
}

func TestMustGet(t *testing.T) {
	Reset()

	v := MustGet("object", map[string]interface{}{"name": "fredrik"})
	assert.Equal(t, map[string]interface{}{"name": "fredrik"}, v.(map[string]interface{}))

	config.Set("object", map[string]interface{}{"name": "fredrik"})

	v = MustGet("object")
	assert.Equal(t, map[string]interface{}{"name": "fredrik"}, v.(map[string]interface{}))
}

func TestTypeInt(t *testing.T) {
	Reset()

	v, err := Int("int")
	assert.NotNil(t, err)
	assert.Equal(t, 0, v)

	config.Set("int", 1)

	v, err = Int("int")
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
}

func TestTypeMustInt(t *testing.T) {
	Reset()

	v := MustInt("int", 12)
	assert.Equal(t, 12, v)

	config.Set("int", 13)

	v = MustInt("int")
	assert.Equal(t, 13, v)
}

func TestTypeString(t *testing.T) {
	Reset()

	v, err := String("name")
	assert.NotNil(t, err)
	assert.Equal(t, "", v)

	config.Set("name", "fredrik")

	v, err = String("name")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestTypeMustString(t *testing.T) {
	Reset()

	v := MustString("name", "test")
	assert.Equal(t, "test", v)

	config.Set("name", "fredrik")

	v = MustString("name")
	assert.Equal(t, "fredrik", v)
}

func TestTypeList(t *testing.T) {
	Reset()

	v, err := List("names")
	assert.NotNil(t, err)
	assert.Equal(t, []string{}, v)

	config.Set("names", []string{"fredrik", "elli"})

	v, err = List("names")
	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestTypeMustList(t *testing.T) {
	Reset()

	v := MustList("names", []string{})
	assert.Equal(t, []string{}, v)

	config.Set("names", []string{"fredrik", "elli"})

	v = MustList("names")
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestTypeUint(t *testing.T) {
	Reset()

	v, err := Uint("uint")
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), v)

	config.Set("uint", uint(1))

	v, err = Uint("uint")
	assert.Nil(t, err)
	assert.Equal(t, uint(1), v)
}

func TestTypeMustUint(t *testing.T) {
	Reset()

	v := MustUint("uint", 2)
	assert.Equal(t, uint(2), v)

	config.Set("uint", uint(1))

	v = MustUint("uint")
	assert.Equal(t, uint(1), v)
}
