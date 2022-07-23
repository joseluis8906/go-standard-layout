package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// NewRouter ...
func NewRouter(opts ...OptionFunc) *chi.Mux {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	r := chi.NewRouter()
	r.Use(Logger)
	r.Use(JSONer)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   conf.allowedOrigins,
		AllowedMethods:   conf.allowedMethods,
		AllowedHeaders:   conf.allowedHeaders,
		ExposedHeaders:   conf.exposedHeaders,
		AllowCredentials: conf.allowCredentials,
		MaxAge:           conf.maxAge,
	}))

	return r
}
