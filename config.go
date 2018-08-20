package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

// Localconfig is a Singleton used to store the configuration loaded from the config.json so that others can use it.
var Localconfig *Configuration

func loadConfig(configpath string) error {
	file, _ := os.Open(configpath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return errors.New("Unable to find the given file")
	}
	Localconfig = &configuration
	return nil
}
