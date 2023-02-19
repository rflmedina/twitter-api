package controllers

import (
	"api/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading request body", err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal("Error unmarshalling request body", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	repository := repository.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created with id %d", userID)))

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
