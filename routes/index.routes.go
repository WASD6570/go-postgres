package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wasd6570/go-postgres/middleware"
)

func Init_server() {
	router := mux.NewRouter()
	router.Use(middleware.HeadersMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"DWF mod 7 GOLANG API\"}"))
	})
	AuthRoutes(router)
	User_routes(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))

}
