package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

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
