package config

import (
	"os"
	"testing"

	"github.com/frozzare/go-assert"
)

func setupEnvTest() {
	Use(NewEnv())
}

func TestEnvBool(t *testing.T) {
	setupEnvTest()
	os.Setenv("BOOL", "true")

	v, err := Bool("bool")

	assert.Nil(t, err)
	assert.True(t, v)
}

func TestEnvFloat(t *testing.T) {
	setupEnvTest()
	os.Setenv("FLOAT", "12.13")

	v, err := Float("float")

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

func TestEnvString(t *testing.T) {
	setupEnvTest()
	os.Setenv("NAME", "fredrik")

	v, err := String("name")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestEnvList(t *testing.T) {
	setupEnvTest()
	os.Setenv("NAMES", "fredrik, elli")

	v, err := List("NAMES")

	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestEnvUint(t *testing.T) {
	setupEnvTest()
	os.Setenv("UINT", "1")

	v, err := Uint("uint")

	assert.Nil(t, err)
	assert.True(t, 1 == v)
}

func TestEnvStringDot(t *testing.T) {
	setupEnvTest()
	os.Setenv("NAME_ENV", "fredrik")

	v, err := String("name.env")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}
