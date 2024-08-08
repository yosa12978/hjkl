package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yosa12978/hjkl/config"
	"github.com/yosa12978/hjkl/data"
	"github.com/yosa12978/hjkl/endpoints"
	"github.com/yosa12978/hjkl/logging"
	"github.com/yosa12978/hjkl/middleware"
)

func Run() error {
	router := http.NewServeMux()
	cfg := config.Read()

	// Connecting to databases
	data.Mariadb(context.TODO())
	data.Redis(context.TODO())

	router.HandleFunc("/api/hello", endpoints.SayHello())
	router.HandleFunc("/api/hello/{name}", endpoints.SayHello())

	router.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	logger := logging.NewJsonLogger(os.Stdout)
	handlerChain := middleware.Chain(
		router,
		middleware.Latency(logger),
		middleware.StripSlash,
	)

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	return http.ListenAndServe(addr, handlerChain)
}
