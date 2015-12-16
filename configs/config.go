package configs

import (
	"os"
	"strconv"

	"errors"
	"io/ioutil"

	"fmt"

	"github.com/BurntSushi/toml"
)

// Instance represents na instance data from the config file
type Instance struct {
	Domain     string `toml:"domain"`
	WindowHour string `toml:"window_hour"`
	Region     string `toml:"region"`
}

// Config struct represents the global configuration from config file
type Config struct {
	Region    string              `toml:"region"`
	Instances map[string]Instance `toml:"instances"`
}

var configObject *Config

// CreateConfig reads file "path" and returns a new configObject if it doesnt exist yet
func CreateConfig(path string) (*Config, error) {
	if configObject != nil {
		return configObject, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	configObject = new(Config)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if _, err = toml.Decode(string(data), configObject); err != nil {
		return nil, err
	}
	for instanceName, instance := range configObject.Instances {
		var hh int
		if hh, err = strconv.Atoi(instance.WindowHour); err != nil || hh > 23 || hh < 0 {
			return nil, fmt.Errorf("%s: failed to parse window_hour", instanceName)
		}
	}
	return configObject, nil
}

// GetConfig returns configObject if it exists, otherwise error
func GetConfig() (*Config, error) {
	if configObject == nil {
		return nil, errors.New("configObject not created")
	}
	return configObject, nil
}

// ClearConfig deletes configObject
func ClearConfig() error {
	if configObject == nil {
		return errors.New("configObject not created")
	}
	configObject = nil
	return nil
}
