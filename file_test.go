package config

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func init() {
	Use(NewFromBytes("json", []byte(`
        {
            "bytes": {
                "bool": true,
                "float": 12.13,
                "int": 12,
                "string": "fredrik",
                "list": ["fredrik", "elli"],
                "uint": 1
            }
        }
    `)))
}

func TestFileBool(t *testing.T) {
	v, err := Bool("bytes.bool")
	assert.Nil(t, err)
	assert.True(t, v)
}

func TestFileFloat(t *testing.T) {
	v, err := Float("bytes.float")
	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestFileInt(t *testing.T) {
	v, err := Int("bytes.int")
	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestFileString(t *testing.T) {
	v, err := String("bytes.string")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestFileList(t *testing.T) {
	v, err := List("bytes.list")
	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestFileUint(t *testing.T) {
	v, err := Uint("bytes.uint")
	assert.Nil(t, err)
	assert.True(t, 1 == v)
}
