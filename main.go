package main

import (
	"github.com/gorilla/mux"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"

)

var userCollection = db().Database(os.Getenv("DB_NAME")).Collection("users") // get collection "users" from db() which returns *mongo.Client

func main() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", createProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", getUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", deleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
