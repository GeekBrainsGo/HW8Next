package server

import (
	"github.com/go-chi/chi"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) routes() { // 1
	s.mux.Route("/", func(r chi.Router) {
		r.Get("/", s.IndexHadle)
		r.Get("/new", s.NewHandle)
		r.Route("/view", func(r chi.Router) {
			r.Get("/{id}", s.ViewHandle)
		})
		r.Route("/edit", func(r chi.Router) {
			r.Get("/{id}", s.EditHandle)
		})
		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(s.config.SwagURL)))
			r.Get("/"+s.config.SwagJSON, s.SwaggerHandle)
			r.Route("/posts", func(r chi.Router) {
				r.Post("/", s.PostHandle)
				r.Delete("/{id}", s.DeleteHandle)
				r.Put("/{id}", s.PutHandle)
			})
		})
	})
}
