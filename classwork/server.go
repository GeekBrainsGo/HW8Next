package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type Server struct {
	lg   *logrus.Logger
	conf *ServerConfig
}

func NewServer(conf *ServerConfig, lg *logrus.Logger) *Server {
	return &Server{
		lg:   lg,
		conf: conf,
	}
}

func (serv *Server) Start() error {
	r := chi.NewMux()
	r.Route("/", func(r chi.Router) {
		r.Get("/", serv.HandleGetIndex)
	})
	return http.ListenAndServe(serv.conf.Addr, r)
}
