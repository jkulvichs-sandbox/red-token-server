package config

import (
	"github.com/jkulvichs-sandbox/red-token/logger"
	"github.com/jkulvichs-sandbox/red-token/storage"
)

// Config - Configuration file structure
type Config struct {
	Logger  logger.ConfLogger   `yaml:"logger"`
	Storage storage.ConfStorage `yaml:"storage"`
}
