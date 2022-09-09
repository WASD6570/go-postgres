package routes

import (
	"github.com/gorilla/mux"
	authControllers "github.com/wasd6570/go-postgres/controllers/auth"
	userController "github.com/wasd6570/go-postgres/controllers/user"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", authControllers.HandleLogin).Methods("POST")
	router.HandleFunc("/register", userController.HandleCreateUser).Methods("POST")
}
