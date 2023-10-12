package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type ConfigData struct {
	LogLevel   string `env:"LOGLEVEL" toml:"log_level"`
	Port       string `env:"PORT" toml:"port"`
	ServerID   string `env:"SERVERID" toml:"server_id"`
	ReadyzFail bool   `toml:"-"`
	LivezFail  bool   `toml:"-"`
}

// Get loads the toml file contents
func (c *ConfigData) Get(name string) error {
	_, err := toml.DecodeFile(name, &c)

	if err != nil {
		return err
	}

	// Iterate over each key and dump if there's an issue.
	// This is needed, because config problems happen all the.damn.time.
	rvalues := reflect.ValueOf(*c)

	configCheck := true

	log.Info("Checking config")
	for i := 0; i < rvalues.NumField(); i++ {
		if rvalues.Field(i).String() == "" {
			fmt.Println("Config file entry issue: " + rvalues.Type().Field(i).Name)
			configCheck = false
		}
	}

	if !configCheck {
		log.Error("Config issue detected. Exiting...")
		os.Exit(0)
	}

	log.Info("Config OK")

	return nil
}

// Empty checks if the config is empty and returns true if it is and false if not
func (c *ConfigData) IsEmpty() bool {
	rvalues := reflect.ValueOf(*c)

	for i := 0; i < rvalues.NumField(); i++ {
		// get name of field from rvalues.Type().Field(i).Name
		if (rvalues.Type().Field(i).Name) == "AWSConfig" || (rvalues.Type().Field(i).Name) == "ClerkClient" {
			continue // because this is a composite entry in the config struct
		}
		if rvalues.Field(i).String() != "" {
			return false
		}
	}

	return true
}
