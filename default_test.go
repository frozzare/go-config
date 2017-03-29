package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestDefaultValue(t *testing.T) {
	assert.Nil(t, defaultValue(nil))

	assert.False(t, defaultValue(false).(bool))
	assert.True(t, defaultValue(false, true).(bool))

	assert.Equal(t, 12.13, defaultValue(12.13).(float64))
	assert.Equal(t, 12.13, defaultValue(0.0, 12.13).(float64))

	assert.Equal(t, nil, defaultValue(nil))
	assert.Equal(t, map[string]interface{}{}, defaultValue(nil, map[string]interface{}{}).(map[string]interface{}))

	assert.Equal(t, 12, defaultValue(12).(int64))
	assert.Equal(t, 12, defaultValue(0, 12).(int64))

	assert.Equal(t, "hello", defaultValue("hello").(string))
	assert.Equal(t, "hello", defaultValue("", "hello").(string))

	assert.Equal(t, []string{"hello"}, defaultValue([]string{"hello"}).([]string))
	assert.Equal(t, []string{"hello"}, defaultValue([]string{}, []string{"hello"}).([]string))

	assert.Equal(t, uint64(12), defaultValue(uint64(12)).(uint64))
	assert.Equal(t, uint64(12), defaultValue(uint64(0), 12).(uint64))
}
