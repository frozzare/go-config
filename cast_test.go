package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestCastBool(t *testing.T) {
	tests := []struct {
		err      bool
		value    interface{}
		expected bool
	}{
		{false, "true", true},
		{false, "false", false},
		{false, true, true},
		{false, false, false},
		{true, nil, false},
	}

	for _, test := range tests {
		v, err := castBool(test.value)

		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, v)
	}
}

func TestCastFloat(t *testing.T) {
	tests := []struct {
		err      bool
		value    interface{}
		expected float64
	}{
		{false, "12.13", 12.13},
		{false, 12.13, 12.13},
		{true, 12, 0.0},
		{true, nil, 0.0},
	}

	for _, test := range tests {
		v, err := castFloat(test.value)

		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, v)
	}
}

func TestCastInt(t *testing.T) {
	tests := []struct {
		err      bool
		value    interface{}
		expected int
	}{
		{false, "12", 12},
		{false, 12, 12},
		{false, 12.13, 12},
		{true, nil, 0},
	}

	for _, test := range tests {
		v, err := castInt(test.value)

		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, v)
	}
}

func TestCastList(t *testing.T) {
	tests := []struct {
		err      bool
		value    interface{}
		expected []string
	}{
		{false, "fredrik, elli", []string{"fredrik", "elli"}},
		{false, []string{"fredrik", "elli"}, []string{"fredrik", "elli"}},
		{false, nil, []string{}},
	}

	for _, test := range tests {
		v, err := castList(test.value)

		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, v)
	}
}

func TestCastString(t *testing.T) {
	tests := []struct {
		err      bool
		value    interface{}
		expected string
	}{
		{false, "12", "12"},
		{true, 12, ""},
		{true, 12.13, ""},
		{true, nil, ""},
	}

	for _, test := range tests {
		v, err := castString(test.value)

		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, v)
	}
}

func TestCastUint(t *testing.T) {
	tests := []struct {
		err      bool
		value    interface{}
		expected uint
	}{
		{false, "1", 1},
		{false, 1, 1},
		{false, 1.13, 1},
		{true, nil, 0},
	}

	for _, test := range tests {
		v, err := castUint(test.value)

		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, v)
	}
}
