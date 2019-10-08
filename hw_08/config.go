package main

import (
	yaml "github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"serv/logger"
)

type Config struct {
	Logger logger.LoggerConfig `yaml:"logger"`
}

func ReadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	conf := Config{}
	_ = yaml.Unmarshal(data, &conf)

	return &conf, nil
}
