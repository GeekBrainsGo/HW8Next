package main

import (
	"log/syslog"
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Syslog bool   `yaml:"syslog"`
	Output string `yaml:"output"`
}

func ConfigureLogger(conf *LoggerConfig) (*logrus.Logger, error) {
	lg := logrus.New()

	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	lg.SetLevel(level)

	if conf.Syslog {
		sysw, err := syslog.New(syslog.LOG_DEBUG, "serv")
		if err != nil {
			return nil, err
		}
		lg.SetOutput(sysw)
	} else {
		if conf.Output != "" {
			f, err := os.Create(conf.Output)
			if err != nil {
				return nil, err
			}
			lg.SetOutput(f)
			// Система сама закроет файл после закрытия приложения
		}
	}

	return lg, nil
}
