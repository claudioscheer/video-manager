package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	// Config file.
	Config config
)

type config struct {
	Port     uint16
	Database struct {
		User     string
		Password string
		DB       string
		Host     string
		Port     uint16
		Timeout  uint8
	} `yaml:"database"`
}

func init() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("A config file must be specified when starting the application.")
		os.Exit(1)
	}
	c, err := loadConfigFile(args[0])
	if err != nil {
		fmt.Println("The following error occurred while loading the config file:")
		panic(err)
	}
	Config = c
}

func loadConfigFile(configFile string) (config, error) {
	config := config{}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return config, errors.New("the config file does not exists")
	}
	data, err := ioutil.ReadFile(configFile)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
