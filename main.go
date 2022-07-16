package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("server starting...")
	router := mux.NewRouter().StrictSlash(true)
	//main pages
	router.HandleFunc("/", all)
	router.HandleFunc("/home", lk)
	router.HandleFunc("/home/{id}", lkpers)
	router.HandleFunc("/exit", exit)
	router.HandleFunc("/ucp", ucp)

	//rest pages
	router.HandleFunc("/rest/auth/ajax/{login}/{password}", authAjax).Methods("GET")
	router.HandleFunc("/rest/auth/{login}/{token}", auth).Methods("GET")
	router.HandleFunc("/rest/reg/{login}/{email}/{password}", reg).Methods("GET")
	router.HandleFunc("/rest/ucp/create/{id}/{name}/{floor}/{age}/{nazi}/{skin}/{country}/{quenta}", UcpCreate).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	/*
		router.HandleFunc("/", all)
		router.HandleFunc("/home", lk)
		router.HandleFunc("/home/{idstate}", lkuser)
		router.HandleFunc("/exit", exit)
		router.HandleFunc("/ucp", ucp)
		router.HandleFunc("/auth/{login}/{token}", auth).Methods("GET")
		router.HandleFunc("/rest/ucp/count/{login}", CountPerson).Methods("GET")
		router.HandleFunc("/rest/ucp/count/id/{id}", CountPersonId).Methods("GET")
		router.HandleFunc("/rest/ucp/findperson/{nickname}", FindPerson).Methods("GET")
		router.HandleFunc("/rest/ucp/findperson/id/{id}", FindPersonId).Methods("GET")
		router.HandleFunc("/rest/ucp/findpersons", FindPersons).Methods("GET")
		router.HandleFunc("/rest/ucp/create/{login}/{nickname}/{floor}/{age}/{nazi}/{skin}/{country}/{quenta}", UcpCreate).Methods("GET")
		router.HandleFunc("/rest/get/user/{criterion}/{value}", GetOneUser).Methods("GET") //return one user
		router.HandleFunc("/rest/get/users", GetAllUsers).Methods("GET")                   //return 100 users
		router.HandleFunc("/rest/auth/{login}/{password}", Auth).Methods("GET")
		router.HandleFunc("/rest/reg/{login}/{email}/{password}", Registration).Methods("GET")
	*/
	log.Fatal(http.ListenAndServe(":8080", router))
}

//logic lk user
//on page get id 1/4
//in cookie find user
//find curent pers
//exec on page
