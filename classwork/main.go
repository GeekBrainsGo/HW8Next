package main

import (
	"flag"
	"fmt"
)

func main() {
	flagConfigPath := flag.String("c", "./config.yaml", "yaml config file path")
	flag.Parse()

	conf, err := ReadConfig(*flagConfigPath)
	if err != nil {
		panic(fmt.Sprintf("can't read config: %s", err))
	}

	lg, err := ConfigureLogger(&conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("can't configure logger: %s", err))
	}

	serv := NewServer(&conf.Server, lg)
	if err := serv.Start(); err != nil {
		lg.WithError(err).Fatal("can't start the server")
	}
}
