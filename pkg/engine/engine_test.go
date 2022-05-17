package engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {

	type sample struct {
		object map[string]interface{}
		json   string
		yml    string
		toml   string
	}

	testSamples := []sample{}

	testSamples = append(testSamples, sample{
		object: map[string]interface{}{"foo": float64(1), "bar": float64(2)},
		json:   `{"foo":1, "bar":2}`,
		yml:    `foo: 1\nbar: 2`,
		toml:   `foo = 1\nbar = 2`,
	})

	testSamples = append(testSamples, sample{
		object: map[string]interface{}{"foo": map[string]interface{}{"net": "true", "bool": false}, "bar": float64(2)},
		json:   `{"foo":{"net": "true", "bool": false}, "bar":2}`,
		yml:    `foo:\n\tnet: 'true'\n\tbool: false\nbar: 2`,
		toml:   `bar = 2\n[foo]\nnet = "true"\nbool = false\n`,
	})

	t.Run("Read from type to Object", func(t *testing.T) {
		for _, testSample := range testSamples {
			json, _ := ToJson(testSample.json)
			yml, _ := ToYml(testSample.yml)
			toml, _ := ToToml(fmt.Sprint(testSample.toml))

			assert.EqualValues(t, testSample.object, json, "The convertion from json to map has a problem")
			assert.EqualValues(t, testSample.object, yml, "The convertion from yml to map has a problem")
			assert.EqualValues(t, testSample.object, toml, "The convertion from toml to map has a problem")
		}
	})

	t.Run("Write Object as type", func(t *testing.T) {

	})

	t.Run("Validate file format", func(t *testing.T) {

	})

	t.Run("Format discovery", func(t *testing.T) {})

}
