package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controllers"
	"./driver"
	"./utils"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectionDB()

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/welcome", utils.TokenVerifyMiddleWare(controller.ProtectedEndpoint())).Methods("GET")

	log.Println("Listen on port 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Innopolis is the best."))
}
