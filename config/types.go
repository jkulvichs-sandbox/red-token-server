package config

import (
	"server/logger"
	"server/storage"
)

// Config - Configuration file structure
type Config struct {
	Logger  logger.ConfLogger   `yaml:"logger"`
	Storage storage.ConfStorage `yaml:"storage"`
}
