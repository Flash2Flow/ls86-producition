package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	bool, err := ls86.Server.checkAll()
	if bool == true {
		ls86.Server.Start()
	} else {
		log.Println(err)
	}
}

func (s *Server) Start() {
	log.Println("server starting...")
	router := mux.NewRouter().StrictSlash(true)
	//main pages
	router.HandleFunc("/", all)
	router.HandleFunc("/exit", exit)
	router.HandleFunc("/home", home)
	router.HandleFunc("/home/{id}", home)
	router.HandleFunc("/auth/{login}/{token}", auth)
	router.HandleFunc("/ucp", ucp)

	//rest pages
	router.HandleFunc("/rest/ucp", ucp).Methods("GET")
	router.HandleFunc("/rest/reg/{login}/{email}/{password}", reg).Methods("GET")
	router.HandleFunc("/rest/auth/{login}/{password}", authRest).Methods("GET")
	router.HandleFunc("/rest/user/{value}", getOne).Methods("GET")
	/*
		router.HandleFunc("/home", lk)
		router.HandleFunc("/home/{id}", lkpers)
		router.HandleFunc("/exit", exit)
		router.HandleFunc("/ucp", ucp)

		//rest pages
		router.HandleFunc("/rest/auth/ajax/{login}/{password}", authAjax).Methods("GET")
		router.HandleFunc("/rest/auth/{login}/{token}", auth).Methods("GET")
		router.HandleFunc("/rest/reg/{login}/{email}/{password}", reg).Methods("GET")
		router.HandleFunc("/rest/ucp/create/{id}/{name}/{floor}/{age}/{nazi}/{skin}/{country}/{quenta}", UcpCreate).Methods("GET")
	*/

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Fatal(http.ListenAndServe(ls86.Server.Port, router))
}
