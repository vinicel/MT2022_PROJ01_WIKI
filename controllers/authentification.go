package controllers

import (
	"encoding/json"
	"github.com/vinicel/MT2022_PROJ01_WIKI/connector"
	"github.com/vinicel/MT2022_PROJ01_WIKI/models"
	u "github.com/vinicel/MT2022_PROJ01_WIKI/utils"
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
		Account: models.Accounts{},
	}
	err = json.Unmarshal(body, &accounts.Account)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	if !strings.Contains(accounts.Account.Email, "@") {
		http.Error(w, "email: Is not an email", 400)
		return
	}
	var user models.Database
	connector.Db.Where("email = ?", accounts.Account.Email).First(&user.Account)
	if user.CheckPassword(accounts.Account.Password) {
		token, err := accounts.GenerateJWT(user.Account.ID, user.Account.Email)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&token)
	} else {
		http.Error(w, "Bad Login or password", 403)
		return
	}
}