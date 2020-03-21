// Package server implements blog sever with mongodb and swagger.
package server

import (
	"context"
	"encoding/json"
	"net/http"
	"remoteblog/models"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Config stands for yaml config struct.
type Config struct {
	Root         string `yaml:"rootdir"`
	Addr         string `yaml:"addr"`
	TemplatesDir string `yaml:"templatesdir"`
	DBname       string `yaml:"dbname"`
	SwagURL      string `yaml:"swagurl"`
	SwagJSON     string `yaml:"swagpath"`
}

// Server stands for server struct.
type Server struct {
	config *Config
	ctx    context.Context
	db     *mongo.Database
	lg     *logrus.Logger
	mux    *chi.Mux
	Page   models.Page
	server *http.Server
}

// New creates new server.
func New(ctx context.Context, lg *logrus.Logger, conf *Config, db *mongo.Database) *Server { // 1
	return &Server{
		ctx:    ctx,
		lg:     lg,
		config: conf,
		mux:    chi.NewMux(),
		db:     db,
		Page: models.Page{
			Title: "Awsome Blog",
		},
	}
}

// Start starts new server.
func (s *Server) Start() *Server { // 2
	s.server = &http.Server{
		Addr:    s.config.Addr,
		Handler: s.mux,
	}
	s.routes()
	go func() {
		s.lg.Info("server started")
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			s.lg.Errorf("server: %s", err)
		}
	}()
	return s
}

// Stop stops the server.
func (s *Server) Stop() error { // 1
	s.lg.Info("stopping server")
	return s.server.Shutdown(s.ctx)
}

// SendErr sends and log error to user.
func (s *Server) SendErr(w http.ResponseWriter, err error, code int, obj ...interface{}) { // 1
	s.lg.WithField("data", obj).WithError(err).Error("server error")
	w.WriteHeader(code)
	errModel := models.ErrorModel{
		Code:     code,
		Err:      err.Error(),
		Desc:     "server error",
		Internal: obj,
	}
	data, _ := json.Marshal(errModel)
	w.Write(data)
}

// SendInternalErr sends 500 error.
func (s *Server) SendInternalErr(w http.ResponseWriter, err error, obj ...interface{}) { // 1
	s.SendErr(w, err, http.StatusInternalServerError, obj)
}
