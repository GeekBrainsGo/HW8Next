package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Logger LoggerConfig `yaml:"logger"`
	Server ServerConfig `yaml:"server"`
}

func ReadConfig(path string) (*Config, error) {
	f, _ := os.Open(path)
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	conf := Config{}
	_ = yaml.Unmarshal(data, &conf)

	return &conf, nil
}
