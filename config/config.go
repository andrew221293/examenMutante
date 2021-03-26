package config

import (
	"encoding/json"
	"io/ioutil"
)

// LoadConfig loads the configuration
// on the config.json file; for safety, this will
// be ignored and will be created in every single environment manually
func LoadConfig(structure interface{}) error {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, structure)
}
