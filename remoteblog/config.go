package main

import (
	"io/ioutil"
	"os"
	"remoteblog/server"

	"gopkg.in/yaml.v2"
)

// Config stands overall app config.
type Config struct {
	Logger LoggerConfig  `yaml:"logger"`
	Server server.Config `yaml:"server"`
}

// ReadConfig reads yaml to app config.
func ReadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
