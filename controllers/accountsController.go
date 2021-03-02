package controllers

import (
	"encoding/json"
	"github.com/Wiki-Go/models"
	u "github.com/Wiki-Go/utils"
	"io/ioutil"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	accounts := models.Accounts{}
	err = json.Unmarshal(body, &accounts)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	resp := accounts.CreateUser()
	u.Respond(w, resp)
}