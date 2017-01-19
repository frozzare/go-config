package config

import (
	"os"
	"testing"

	"github.com/frozzare/go-assert"
)

func init() {
	Use(NewEnv())
}

func TestEnvBool(t *testing.T) {
	os.Setenv("BOOLENV", "true")
	v, err := Bool("boolenv")
	assert.Nil(t, err)
	assert.True(t, v)
}

func TestEnvFloat(t *testing.T) {
	os.Setenv("FLOATENV", "12.13")
	v, err := Float("floatenv")
	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestEnvInt(t *testing.T) {
	os.Setenv("INTENV", "12")
	v, err := Int("intenv")
	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestEnvString(t *testing.T) {
	os.Setenv("NAMEENV", "fredrik")

	v, err := String("nameenv")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestEnvList(t *testing.T) {
	os.Setenv("NAMESENV", "fredrik, elli")
	v, err := List("namesenv")
	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestEnvUint(t *testing.T) {
	os.Setenv("UINTENV", "1")
	v, err := Uint("uint")
	assert.Nil(t, err)
	assert.True(t, 1 == v)
}

func TestEnvStringDot(t *testing.T) {
	os.Setenv("NAME_ENV", "fredrik")

	v, err := String("name.env")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}
