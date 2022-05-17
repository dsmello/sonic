package engine

import (
	"encoding/json"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// Responsible to read from (Json, yml, toml) and convert to map[string]interface{}

func ToJson(obj string) (map[string]interface{}, error) {
	x := map[string]interface{}{}

	err := json.Unmarshal([]byte(obj), &x)
	return x, err
}

func ToYml(obj string) (map[string]interface{}, error) {
	x := map[string]interface{}{}

	err := yaml.Unmarshal([]byte(obj), &x)
	return x, err
}

func ToToml(obj string) (map[string]interface{}, error) {
	x := map[string]interface{}{}

	err := toml.Unmarshal([]byte(obj), &x)
	return x, err

}
