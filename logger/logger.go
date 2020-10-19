package logger

import "github.com/sirupsen/logrus"

// ConfLogger - Logger configuration
type ConfLogger struct {
	// Formatter - Log formatting: json, text
	Formatter string `yaml:"formatter"`
	// Level - Minimal log level: panic, fatal, error, warn, info, debug, trace
	Level string `yaml:"level"`
	// ReportCaller - Show or not trace for log function
	ReportCaller bool `yaml:"report"`
}

// New - Configures app logger
func New(conf ConfLogger) (*logrus.Logger, error) {
	logger := logrus.New()

	// Setting formatter
	if conf.Formatter == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	// Setting log level
	lvl, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(lvl)

	// Setting report caller
	logger.SetReportCaller(conf.ReportCaller)

	return logger, nil
}
