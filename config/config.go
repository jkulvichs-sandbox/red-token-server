package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"github.com/jkulvichs-sandbox/red-token/logger"
	"github.com/jkulvichs-sandbox/red-token/storage"
)

// GenerateConfig - Makes default config in path
func GenerateConfig(path string) error {
	conf := Config{
		Logger: logger.ConfLogger{
			Formatter:    "text",
			Level:        "info",
			ReportCaller: false,
		},
		Storage: storage.ConfStorage{
			SQLite: &storage.ConfSQLiteDataBase{
				Path: "db.sqlite",
			},
		},
	}

	confBytes, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, confBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// LoadConfig - Loads config from YAML file
func LoadConfig(path string, createIfNotExist ...bool) (*Config, error) {
	createDefault := false
	if len(createIfNotExist) > 0 {
		createDefault = createIfNotExist[0]
	}

	// Checking for config or creating default one
	if _, err := os.Open(path); err != nil {
		if os.IsNotExist(err) && createDefault {
			if err := GenerateConfig(path); err != nil {
				return nil, fmt.Errorf("can't create default config: %s", err)
			}
		} else {
			return nil, fmt.Errorf("config doesn't exist")
		}
	}

	// Reading config
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := yaml.Unmarshal(confBytes, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// StoreConfig - Stores config into file
func StoreConfig(path string, conf *Config) error {
	confBytes, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, confBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
