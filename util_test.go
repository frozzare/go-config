package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestValue(t *testing.T) {
	v, err := value("name", nil)
	assert.NotNil(t, err)
	assert.Nil(t, v)

	v, err = value("name", map[string]interface{}{})
	assert.NotNil(t, err)
	assert.Nil(t, v)

	v, err = value("name", map[string]interface{}{"name": "fredrik"})
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)

	v, err = value("user.name", map[string]interface{}{"user": map[string]interface{}{"name": "fredrik"}})
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)

	v, err = value("user.ok", map[string]interface{}{"user": map[string]interface{}{"ok": true}})
	assert.Nil(t, err)
	assert.Equal(t, true, v)

	v, err = value("user.number", map[string]interface{}{"user": map[string]interface{}{"number": 42}})
	assert.Nil(t, err)
	assert.Equal(t, 42, v)

	v, err = value("name", "Fredrik")
	assert.Nil(t, err)
	assert.Equal(t, "Fredrik", v)
}

func TestToString(t *testing.T) {
	assert.Equal(t, "fredrik", toString("fredrik"))
	assert.Equal(t, "1", toString(1))
	assert.Equal(t, "true", toString(true))
	assert.Equal(t, "1.3", toString(1.3))
	assert.Equal(t, "map[]", toString(map[string]interface{}{}))
	assert.Equal(t, "<nil>", toString(nil))
}
