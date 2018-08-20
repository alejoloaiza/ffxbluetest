package main

import "encoding/json"
import "os"
import "fmt"

var Localconfig *Configuration

func loadConfig(configpath string) {
	file, _ := os.Open(configpath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	Localconfig = &configuration
}
