package config

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/frozzare/go-assert"
)

func setupFileTest() {
	Reset()

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
	setupFileTest()

	v, err := Bool("bytes.bool")

	assert.Nil(t, err)
	assert.True(t, v)
}

func TestFileFloat(t *testing.T) {
	setupFileTest()

	v, err := Float("bytes.float")

	assert.Nil(t, err)
	assert.Equal(t, 12.13, v)
}

func TestFileInt(t *testing.T) {
	setupFileTest()

	v, err := Int("bytes.int")

	assert.Nil(t, err)
	assert.Equal(t, 12, v)
}

func TestFileString(t *testing.T) {
	setupFileTest()

	v, err := String("bytes.string")

	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)
}

func TestFileList(t *testing.T) {
	setupFileTest()

	v, err := List("bytes.list")

	assert.Nil(t, err)
	assert.Equal(t, []string{"fredrik", "elli"}, v)
}

func TestFileUint(t *testing.T) {
	setupFileTest()

	v, err := Uint("bytes.uint")

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), v)
}

func TestWatchFile(t *testing.T) {
	Reset()

	err := exec.Command("cp", "data/config.json", "/tmp/config-watch.json").Run()
	assert.Nil(t, err)

	Use(NewFromFile("/tmp/config-watch.json", true))

	v, err := String("name")
	assert.Nil(t, err)
	assert.Equal(t, "fredrik", v)

	err = exec.Command("cp", "data/config2.json", "/tmp/config-watch.json").Run()
	assert.Nil(t, err)

	time.Sleep(60)

	v, err = String("name")
	assert.Nil(t, err)
	assert.Equal(t, "go", v)

	os.Remove("/tmp/config-watch.json")
}
