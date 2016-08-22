package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Twitter is the name of the Twitter module.
const Twitter string = "twitter"

// Config contains the application's configuration.
var Config Configuration

// Configuration is a structure that contains a slice of modules.
type Configuration struct {
	Modules []Module // A slice of modules.
}

// Module is a structure that contains the name of a module, and its payload.
type Module struct {
	Name    string // The name of the module.
	Payload string // Base64-encoded payload, which will be parsed as JSON.
}

// LoadConfig loads the config from the default location 'config.json'.
func LoadConfig() error {
	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		return err
	}

	return json.Unmarshal(data, &Config)
}

// GetPayload gets the payload for the specified module name.
func GetPayload(moduleName string) (string, error) {
	for _, m := range Config.Modules {
		if m.Name == moduleName {
			return m.Payload, nil
		}
	}

	return "", fmt.Errorf("No payload found for module %s.", moduleName)
}
