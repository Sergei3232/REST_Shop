package controllers

import (
	"encoding/json"
	"github.com/Sergei3232/REST_Shop/models"
	u "github.com/Sergei3232/REST_Shop/utils"
	"net/http"
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request

	if n := models.RoleAvailable(user, "manager"); !n {
		u.Respond(w, u.Message(403, "The method is not available"))
		return
	}

	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		u.Respond(w, u.Message(400, "Error while decoding request body"))
		return
	}

	resp := product.Create()
	u.Respond(w, resp)
}
