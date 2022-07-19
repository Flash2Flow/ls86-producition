package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type REST struct {
	Public  Public
	Private Private
	Secret  Secret
}

type Public struct {
	User    User
	Error   Error
	Success Success
	/*
		public rest logic
		only registration + auth + get one user
		get uuid, and checking actions
		if actions if bad, block uuid
	*/
}

type UserPublic struct {
	Id    int
	Login string
	Email string
}

type Private struct {
	/*
		private rest logic
		only actions who have values him
	*/
}

type Secret struct {
	/*
		secret rest logic
		all actions with hard logics + permissions
		if 2/2 trying register this rest block all cookies + database params
	*/
}

type Success struct{}

func (p Public) Auth() {

}

func (u User) Reg(restType, login, password, email string, w http.ResponseWriter) {
	switch restType {
	case "public":
		use, err := ls86.Data.User.FindOne("login", login)
		if use.Id == 0 {
			switch err {
			case ls86.Data.Error.NotFound:
				e, err := ls86.Data.User.FindOne("email", email)
				if e.Id == 0 {
					switch err {
					case ls86.Data.Error.NotFound:
						pass, err := HashPassword(password)
						if err != nil {
							log.Println("CRITICAL ERROR HASH PASSWORD!!!")
							log.Println(err)
						} else {
							token := uuid.New()
							user := User{
								Login:     login,
								Email:     email,
								Password:  pass,
								AuthToken: token.String(),
							}
							ls86.Data.User.Create(&user)
							ls86.REST.Public.Success.Created(user)
							w.WriteHeader(http.StatusOK)
							json.NewEncoder(w).Encode(user)
						}

					}
				} else {
					ls86.REST.Public.Error.UserEmailAlreadyUsed(email)
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				}

			}
		} else {
			ls86.REST.Public.Error.UserLoginAlreadyUsed(login)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}

	case "private":

	case "secret":

	}

}

func (u User) Get(value string, restType string, title string, w http.ResponseWriter) {

	switch restType {
	case "public":
		u, err := ls86.Data.User.FindOne(title, value)
		switch err {
		case ls86.Data.Error.NotFound:
			ls86.REST.Public.Error.NotFoundReq(value)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, ls86.Data.Error.NotFound.Error())
		case nil:
			userPublic := UserPublic{
				Id:    u.Id,
				Login: u.Login,
				Email: u.Email,
			}
			ls86.REST.Public.Success.Found(userPublic)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(userPublic)
		default:
			ls86.REST.Public.Error.StatusBadRequest(value)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
		}
	case "private":

	case "secret":

	}

}

func (u User) Gets(value string, restType string, title string) {

	switch restType {
	case "public":

	case "private":

	case "secret":

	default:
		//err

	}

}

func reg(w http.ResponseWriter, r *http.Request) {
	//init
	params := mux.Vars(r)
	//url params
	login := params["login"]
	password := params["password"]
	email := params["email"]
	//headers
	restType := r.Header.Get("restType")
	log.Println()

	ls86.REST.Public.User.Reg(restType, login, password, email, w)
}

//pages
func getOne(w http.ResponseWriter, r *http.Request) {
	//init
	params := mux.Vars(r)
	//url params
	value := params["value"]
	//headers
	restType := r.Header.Get("restType")
	title := r.Header.Get("title")

	ls86.REST.Public.User.Get(value, restType, title, w)
}
