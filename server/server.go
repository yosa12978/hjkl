package server

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yosa12978/hjkl/endpoints"
	"github.com/yosa12978/hjkl/middleware"
)

func New(opts ...optionFunc) http.Handler {
	options := newOptions(opts...)
	router := http.NewServeMux()
	addRoutes(router, options)
	var handler http.Handler = router
	handler = middleware.Pipeline(
		router,
		middleware.Latency(options.logger),
		middleware.StripSlash,
		middleware.Recovery(options.logger),
	)
	return handler
}

func addRoutes(r *http.ServeMux, options options) {
	r.HandleFunc("GET /health", endpoints.Health())
	r.HandleFunc("GET /api/hello", endpoints.SayHello())
	r.HandleFunc("GET /api/hello/{name}", endpoints.SayHello())
	r.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
