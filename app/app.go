package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/yosa12978/hjkl/config"
	"github.com/yosa12978/hjkl/data"
	"github.com/yosa12978/hjkl/logging"
	"github.com/yosa12978/hjkl/middleware"
	"github.com/yosa12978/hjkl/util"
)

func Run() error {
	router := http.NewServeMux()
	cfg := config.Read()

	// Connecting to databases
	data.Mariadb(context.TODO())
	data.Redis(context.TODO())

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		util.WriteJson(w, 200, map[string]any{"message": "Hello hjkl!"})
	})

	router.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

	logger := logging.NewJsonLogger(os.Stdout)
	handlerChain := middleware.Chain(
		router,
		middleware.Latency(logger),
	)

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	return http.ListenAndServe(addr, handlerChain)
}
