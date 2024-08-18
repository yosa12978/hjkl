package server

import (
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yosa12978/hjkl/endpoints"
	"github.com/yosa12978/hjkl/logging"
	"github.com/yosa12978/hjkl/middleware"
)

func NewServer() http.Handler {
	router := http.NewServeMux()
	logger := logging.NewJsonLogger(os.Stdout)
	addRoutes(router, logger)
	var handler http.Handler = router
	handler = middleware.Chain(
		router,
		middleware.Latency(logger),
		middleware.StripSlash,
	)
	return handler
}

func addRoutes(r *http.ServeMux, logger logging.Logger) {
	r.HandleFunc("GET /health", endpoints.Health())
	r.HandleFunc("GET /api/hello", endpoints.SayHello())
	r.HandleFunc("GET /api/hello/{name}", endpoints.SayHello())
	r.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
