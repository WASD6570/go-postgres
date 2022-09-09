package userController

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/wasd6570/go-postgres/db"
	"github.com/wasd6570/go-postgres/models"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var User models.User

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\": \"Invalid request\"}"))
		return
	}

	_, error := CreateUser(&User)
	if error != nil {
		log.Print(error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Error creating user\"}"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(User)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	params := mux.Vars(r)
	db.Conn.First(&User).Where("id = ?", params["id"])
	if User.ID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\": \"User not found\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var reqUser models.User
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&reqUser)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\": \"Invalid request\"}"))
		return
	}

	var User models.User

	db.Conn.Model(&User).Where("id = ?", params["id"]).UpdateColumns(models.User{Name: reqUser.Name, LastName: reqUser.LastName})

	if User.ID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\": \"User not found\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}
