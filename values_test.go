package config

import (
	"testing"

	assert "github.com/frozzare/go-assert"
)

func setupValuesTest() {
	Reset()

	Use(NewFromValues(map[string]interface{}{
		"bool":  true,
		"float": 12.13,
		"int":   12,
		"name":  "fredrik",
		"names": []string{"fredrik", "elli"},
		"uint":  1,
		"object": map[string]interface{}{
			"name": "fredrik",
		},
	}))
}

func TestValuesBool(t *testing.T) {
	setupValuesTest()

	v, err := Bool("bool")

	assert.Nil(t, err)
	assert.True(t, v)
}

func TestValuesFloat(t *testing.T) {
	setupValuesTest()
	v, err := Float("float")

	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestValuesInt(t *testing.T) {
	setupValuesTest()
	v, err := Int("int")

	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestValuesString(t *testing.T) {
	setupValuesTest()
	v, err := String("name")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestValuesList(t *testing.T) {
	setupValuesTest()
	v, err := List("names")

	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestValuesUint(t *testing.T) {
	setupValuesTest()
	v, err := Uint("uint")

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), v)
}

func TestValuesStringDot(t *testing.T) {
	setupValuesTest()
	v, err := String("object.name")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}
