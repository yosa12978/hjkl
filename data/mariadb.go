package data

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yosa12978/hjkl/config"
)

var (
	mdb     *sql.DB
	mdbOnce sync.Once
)

func Mariadb(ctx context.Context) *sql.DB {
	mdbOnce.Do(func() {
		cfg := config.Read()
		addr := fmt.Sprintf(
			"%s:%s@%s/%s?multiStatements=true",
			cfg.Mariadb.Username,
			cfg.Mariadb.Password,
			cfg.Mariadb.Addr,
			cfg.Mariadb.Db,
		)
		conn, err := sql.Open("mysql", addr)
		if err != nil {
			panic(err)
		}
		if err := conn.PingContext(ctx); err != nil {
			panic(err)
		}
	})
	return mdb
}
