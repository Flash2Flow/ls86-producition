package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	logs.Logging(logs.Server.Starting)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", all)
	router.HandleFunc("/home", lk)
	router.HandleFunc("/home/{user}", lkuser)
	router.HandleFunc("/exit", exit)
	router.HandleFunc("/auth/{login}/{token}", auth).Methods("GET")
	//token rest must be in headers
	router.HandleFunc("/rest/get/user/{criterion}/{value}", GetOneUser).Methods("GET") //return one user
	router.HandleFunc("/rest/get/users", GetAllUsers).Methods("GET")                   //return 100 users
	router.HandleFunc("/rest/auth/{login}/{password}", Auth).Methods("GET")
	router.HandleFunc("/rest/reg/{login}/{email}/{password}", Registration).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Fatal(http.ListenAndServe(logs.Server.Port, router))
}
