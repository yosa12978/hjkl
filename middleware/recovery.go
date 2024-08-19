package middleware

import (
	"fmt"
	"net/http"

	"github.com/yosa12978/hjkl/logging"
	"github.com/yosa12978/hjkl/util"
)

func Recovery(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(fmt.Sprintf("%+v", err))
					util.WriteMsg(w,
						http.StatusInternalServerError,
						"Internal Server Error",
					)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
