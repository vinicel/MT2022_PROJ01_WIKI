package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wiki/models"
	u "wiki/utils"
)


func (c *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	accounts := models.Database{
		Db: c.Db,
		Account: models.Accounts{},
	}
	err = json.Unmarshal(body, &accounts.Account)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	resp := accounts.CreateUser()
	u.Respond(w, resp)
}