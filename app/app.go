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

	router.HandleFunc("/api/hello", hello())
	router.HandleFunc("/api/hello/{name}", hello())

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

// hello handler godoc
//	@Summary		Say hello
//	@Description	just returns hello message
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/hello [get]
func hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			name = "hjkl"
		}
		util.WriteJson(w, 200,
			map[string]any{
				"message": fmt.Sprintf("Hello %s!", name),
			},
		)
	}
}
