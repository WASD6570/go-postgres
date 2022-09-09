package middleware

import (
	"net/http"

	apiUtils "github.com/wasd6570/go-postgres/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, err := apiUtils.DecodeToken(r.Header.Get("Authorization"))

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("{\"message\": \"Unauthorized\"}"))
			return
		}

		r.Header.Set("user_id", user.ID.String())

		next.ServeHTTP(w, r)
	})
}
