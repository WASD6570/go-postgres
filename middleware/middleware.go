package middleware

import "net/http"

func Middleware(h http.Handler, middleware ...func(w http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
