package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yosa12978/hjkl/config"
	"github.com/yosa12978/hjkl/data"
	"github.com/yosa12978/hjkl/util"
)

func Run() error {
	router := http.NewServeMux()
	cfg := config.Read()

	// Connecting to databases
	data.Mariadb(context.TODO())
	data.Redis(context.TODO())

	router.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		util.WriteJson(w, 200, map[string]any{"message": "Hello hjkl!"})
	})

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	return http.ListenAndServe(addr, router)
}
