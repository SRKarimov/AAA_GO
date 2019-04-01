package userRepository

import (
	"database/sql"
	"log"

	"../../models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	stmt := "INSERT INTO users (email, password, ip) VALUES ($1, $2, $3) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password, user.Ip).Scan(&user.Id)
	logFatal(err)

	user.Password = ""
	return user
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Ip)

	if err != nil {
		return user, err
	}

	return user, nil
}
