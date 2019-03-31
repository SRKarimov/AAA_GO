package models

type User struct {
	Id       int    `json:"Id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
