package authControllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/wasd6570/go-postgres/models"
	apiUtils "github.com/wasd6570/go-postgres/utils"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqAuth models.Auth

	err := json.NewDecoder(r.Body).Decode(&reqAuth)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\": \"Invalid request\"}"))
		return
	}

	Auth := FindOneAuth(&reqAuth)

	if Auth.ID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\": \"User not found\"}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Auth)

}

func HandleRegister(w http.ResponseWriter, r *http.Response) {
	var auth models.Auth
	var user models.User
	var body map[string]interface{}

	defer apiUtils.PanicHandler(w)

	json.NewDecoder(r.Body).Decode(&body)

	for _, m := range body {
		fmt.Printf("m: %v\n", m)
	}

	email := body["email"].(string)
	password := body["password"].(string)
	name := body["name"].(string)
	lastName := body["lastName"].(string)

	user.Name = name
	user.LastName = lastName
	auth.Email = email
	auth.Password = password

	createdUser := CreateUserAndAuth(&auth, &user)

	token, err := apiUtils.EncodeToken(createdUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Error creating user\"}"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("{\"token\": \"" + token + "\"}")
}
