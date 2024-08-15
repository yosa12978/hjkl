package data

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/yosa12978/hjkl/config"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func Postgres(ctx context.Context) *sql.DB {
	dbOnce.Do(func() {
		cfg := config.Read()
		addr := fmt.Sprintf(
			"postgres://%s:%s@%s",
			cfg.Postgres.Username,
			cfg.Postgres.Password,
			cfg.Postgres.Addr,
		)
		conn, err := sql.Open("postgres", addr)
		if err != nil {
			panic(err)
		}
		if err := conn.PingContext(ctx); err != nil {
			panic(err)
		}
		db = conn
	})
	return db
}
