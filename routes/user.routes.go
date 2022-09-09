package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/wasd6570/go-postgres/controllers/user"
	"github.com/wasd6570/go-postgres/middleware"
)

func User_routes(router *mux.Router) {
	router.Handle("/users/{id}", middleware.Middleware(
		http.HandlerFunc(controllers.HandleGetUser),
		middleware.AuthMiddleware)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.HandleUpdateUser).Methods("PATCH")
}
