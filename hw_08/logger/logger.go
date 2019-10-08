package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerConfig struct {
	Level   string `yaml:"level"`
	Output  string `yaml:"output"`
	Rootdir string `yaml:"rootdir"`
	Addr    string `yaml:"addr"`
}

func ConfigureLogger(conf *LoggerConfig) (*logrus.Logger, error) {
	lg := logrus.New()

	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})

	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	lg.SetLevel(level)

	if conf.Output != "" {
		f, _ := os.Create(conf.Output)
		if err != nil {
			return nil, err
		}
		lg.SetOutput(f)
	}

	return lg, nil
}
