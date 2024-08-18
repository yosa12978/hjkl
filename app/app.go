package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yosa12978/hjkl/config"
	"github.com/yosa12978/hjkl/data"
	"github.com/yosa12978/hjkl/server"
)

func Run() error {
	cfg := config.Read()

	// Initializing database connections
	data.Postgres(context.TODO())
	data.Redis(context.TODO())

	srv := server.NewServer()

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	return http.ListenAndServe(addr, srv)
}
