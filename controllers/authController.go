package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Sergei3232/REST_Shop/models"
	u "github.com/Sergei3232/REST_Shop/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	fmt.Println(r.Body)
	if err != nil {
		u.Respond(w, u.Message(400, "Invalid request"))
		return
	}

	resp := account.Create() //Создать аккаунт
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Users{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(400, "Invalid request"))
		return
	}

	resp := models.Login(account.Account, account.Password)
	u.Respond(w, resp)
}
