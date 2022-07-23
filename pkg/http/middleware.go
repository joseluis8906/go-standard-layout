package http

import (
	"net/http"

	"github.com/joseluis8906/go-standard-layout/pkg/log"
	"github.com/urfave/negroni"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Logger().Infof("<-- [%s] %s", r.Method, r.RequestURI)
		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.Status()
		log.Logger().Infof("--> [%s] %s [%d %s]", r.Method, r.RequestURI, statusCode, http.StatusText(statusCode))
	})
}

func JSONer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
