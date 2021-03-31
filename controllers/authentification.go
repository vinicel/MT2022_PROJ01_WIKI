package controllers

import (
	"encoding/json"
	"github.com/Wiki-Go/models"
	u "github.com/Wiki-Go/utils"
	customHTTP "github.com/hellojebus/go-envoz-api/http"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request)  {
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
	if !strings.Contains(accounts.Account.Email, "@") {
		http.Error(w, err.Error(), 500)
		return
	}
	if !strings.Contains(accounts.Account.Email, "@") {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Ceci n'est pas un email !")
	}

	var user models.Database
	accounts.Db.Where("email = ?", accounts.Account.Email).First(&user.Account)
	if user.CheckPassword(accounts.Account.Password) {
		token, err := accounts.GenerateJWT(user.Account.ID, user.Account.Email)
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
			return
		}
		json.NewEncoder(w).Encode(&token)
	} else {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Password incorrect")
		return
	}
}