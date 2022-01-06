package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Check(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	repository := repositories.NewUserRepository(db)
	user, err = repository.Create(user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := r.URL.Query().Get("user")

	db := database.GetDB()

	repository := repositories.NewUserRepository(db)
	users, err := repository.List(nameOrNick)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário pelo Id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando suário"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário"))
}
