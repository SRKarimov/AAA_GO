package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	userRepository "../repository/user"
	u "../utils"

	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing."
			u.RespondWithError(w, http.StatusBadRequest, error.Message)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing."
			u.RespondWithError(w, http.StatusBadRequest, error.Message)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal(err)
		}

		user.Password = string(hash)
		user.Ip = r.RemoteAddr

		userRepo := userRepository.UserRepository{}
		user = userRepo.Signup(db, user)

		if err != nil {
			fmt.Println(err)
			error.Message = "Server error"
			u.RespondWithError(w, http.StatusInternalServerError, error.Message)
			return
		}

		user.Password = ""
		w.Header().Set("Content-Type", "application/json")
		u.ResponseJSON(w, user)
	}
}

func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var jwt models.JWT
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing."
			u.RespondWithError(w, http.StatusBadRequest, error.Message)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing."
			u.RespondWithError(w, http.StatusBadRequest, error.Message)
			return
		}

		userRepo := userRepository.UserRepository{}
		user, err := userRepo.Login(db, user)

		password := user.Password

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				u.RespondWithError(w, http.StatusBadRequest, error.Message)
				return
			} else {
				log.Fatal(err)
			}
		}

		hashedPassword := user.Password

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {
			error.Message = "Invalid Password"
			u.RespondWithError(w, http.StatusUnauthorized, error.Message)
			return
		}

		token, err := u.GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token

		u.ResponseJSON(w, jwt)
	}
}
