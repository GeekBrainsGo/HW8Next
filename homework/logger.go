package main

import "github.com/sirupsen/logrus"

//NewLogger - Creates and returns new logger
func NewLogger(conf LoggerConfig) (error, *logrus.Logger) { //2
	lg := logrus.New()
	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})
	level, err := logrus.ParseLevel(conf.DebugLevel)
	if err != nil {
		return err, nil
	}
	lg.SetLevel(level)
	return nil, lg
}
