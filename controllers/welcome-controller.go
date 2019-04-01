package controllers

import "net/http"
import u "../utils"

func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u.ResponseJSON(w, "Welcome to Innopolis.")
	}
}
