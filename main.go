package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	pgUrl, err := pq.ParseURL("postgres://hell:@localhost/aaa_db?sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/welcome", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Listen on port 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Innopolis is the best."))
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup invoked.")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login invoked.")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProtectedEndpoint invoked.")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare invoked.")
	return nil
}
