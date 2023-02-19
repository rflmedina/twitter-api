package controllers

import (
	"api/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var isARegister bool = true

	if err = user.Prepare(isARegister); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID = uint64(user.ID)

	response.JSON(w, http.StatusCreated, "User created successfully with ID: "+fmt.Sprint(userID))

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUsersRepository(db)
	users, err := repository.SearchByNameOrNick(nameOrNick)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUsersRepository(db)
	user, err := repository.SearchByID(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var isARegister bool = false

	if err = user.Prepare(isARegister); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if err = repository.Update(userID, user); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
