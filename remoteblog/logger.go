package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// LoggerConfig stands for logger config struct.
type LoggerConfig struct {
	Level  string `yaml:"level"`
	Syslog bool   `yaml:"syslog"`
	Output string `yaml:"output"`
}

// ConfigureLogger sets logger.
func ConfigureLogger(lc *LoggerConfig) (*logrus.Logger, error) { // 4
	lg := logrus.New()

	level, err := logrus.ParseLevel(lc.Level)
	if err != nil {
		return nil, err
	}
	lg.SetLevel(level)

	if len(lc.Output) > 0 {
		f, err := os.Create(lc.Output)
		if err != nil {
			return nil, err
		}
		lg.SetOutput(f)
	}
	return lg, nil
}
