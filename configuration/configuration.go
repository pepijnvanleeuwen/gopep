// Copyright (c) 2016 Pepijn van Leeuwen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const Twitter string = "Twitter"

var Config Configuration

type Configuration struct {
	Modules []Module
}

type Module struct {
	Name    string
	Payload Payload
}

type Payload struct {
	Data string // JSON encoded payload.
}

func LoadConfig() error {
	// Try to load the config from the relative path of current executable.
	if d, err := ioutil.ReadFile("config.json"); err != nil {
		return err
	} else {
		return json.Unmarshal(d, &Config)
	}
}

func GetPayload(moduleName string) (Payload, error) {
	for _, m := range Config.Modules {
		if m.Name == moduleName {
			return m.Payload, nil
		}
	}

	return Payload{}, fmt.Errorf("No payload found for module %s.", moduleName)
}
