package config

import (
	"os"
	"testing"

	"github.com/frozzare/go-assert"
)

func setupEnvTest() {
	Reset()

	Use(NewEnv())
}

func TestEnvBool(t *testing.T) {
	setupEnvTest()
	os.Setenv("BOOL", "true")

	v, err := Bool("bool")

	assert.Nil(t, err)
	assert.True(t, v)
}

func TestEnvBoolDefault(t *testing.T) {
	v, err := Bool("booldefault", true)
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}

func TestEnvFloat(t *testing.T) {
	setupEnvTest()
	os.Setenv("FLOAT", "12.13")

	v, err := Float("float")

	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestEnvFloatDefault(t *testing.T) {
	v, err := Float("floatdefault", 12.13)
	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestEnvInt(t *testing.T) {
	setupEnvTest()
	os.Setenv("INT", "12")

	v, err := Int("int")

	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestEnvIntDefault(t *testing.T) {
	v, err := Int("intdefault", 12)
	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestEnvString(t *testing.T) {
	setupEnvTest()
	os.Setenv("NAME", "fredrik")

	v, err := String("name")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestEnvStringDefault(t *testing.T) {
	v, err := String("stringdefault", "hello")
	assert.Nil(t, err)
	assert.Equal(t, "hello", v)
}

func TestEnvList(t *testing.T) {
	setupEnvTest()
	os.Setenv("NAMES", "fredrik, elli")

	v, err := List("NAMES")

	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestEnvListDefault(t *testing.T) {
	v, err := List("listdefault", []string{"hello"})
	assert.Nil(t, err)
	assert.Equal(t, []string{"hello"}, v)
}

func TestEnvUint(t *testing.T) {
	setupEnvTest()
	os.Setenv("UINT", "1")

	v, err := Uint("uint")

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), v)
}

func TestEnvUintDefault(t *testing.T) {
	v, err := Uint("uintdefault", 1)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), v)
}

func TestEnvStringDot(t *testing.T) {
	setupEnvTest()
	os.Setenv("NAME_ENV", "fredrik")

	v, err := String("name.env")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}
