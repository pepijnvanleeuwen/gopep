package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Twitter is the name of the Twitter module.
const Twitter string = "Twitter"

// Config contains the application's configuration.
var Config Configuration

// Configuration is a structure that contains a slice of modules.
type Configuration struct {
	Modules []Module
}

// Module is a structure that contains the name of a module, and its payload.
type Module struct {
	Name    string
	Payload Payload
}

// Payload is a structure that contains JSON-encoded values that can be used
// in the module it belongs to.
type Payload struct {
	Data string // JSON encoded payload.
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
func GetPayload(moduleName string) (Payload, error) {
	for _, m := range Config.Modules {
		if m.Name == moduleName {
			return m.Payload, nil
		}
	}

	return Payload{}, fmt.Errorf("No payload found for module %s.", moduleName)
}
