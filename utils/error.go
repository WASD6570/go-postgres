package apiUtils

import (
	"log"
	"net/http"
)

func PanicHandler(w http.ResponseWriter) {
	rec := recover()
	if rec != nil {
		log.Fatalf("Recovered from panic: %v", rec)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Error creating user\"}"))
	}
}
