package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
)

func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		json.NewDecoder(r.Body).Decode(&user)
		u.ResponseJSON(w, "Welcome "+user.Email+" to Innopolis.")
	}
}
