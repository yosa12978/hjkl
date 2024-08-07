package middleware

import "net/http"

func Chain(f http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
