package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type Rest struct {
	Token string
	Url   Url
}

type Url struct {
	BaseUrl      string
	GetUser      string
	GetUsers     string
	Auth         string
	Registration string
}

var (
	rest = Rest{
		Token: "537003",
		Url: Url{
			BaseUrl:      "http://localhost:" + logs.Server.Port,
			GetUser:      "/rest/get/user/",
			GetUsers:     "/rest/get/users/",
			Auth:         "/rest/auth/",
			Registration: "/rest/reg/",
		},
	}
)

func Auth(w http.ResponseWriter, r *http.Request) {
	//init block
	token := r.Header.Get("rest-token")
	params := mux.Vars(r)
	login := params["login"]
	password := params["password"]

	if token == rest.Token {
		u, err := data.Auth(login, password)
		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(u)
			logs.UserFound(u)
		} else {
			log.Println(customErr.ErrPasswordWrong.Error())
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, customErr.ErrPasswordWrong.Error())
		}

	} else {
		log.Println(customErr.ErrAccessDenied.Error())
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, customErr.ErrAccessDenied.Error())
	}
}

func Registration(w http.ResponseWriter, r *http.Request) {
	//init block
	params := mux.Vars(r)
	login := params["login"]
	email := params["email"]
	password := params["password"]
	token := r.Header.Get("rest-token")
	if token == rest.Token {
		ue, err := data.FindOne("login", login)
		if err == customErr.ErrNotFound {
			_, err := data.FindOne("email", email)
			if err == customErr.ErrNotFound {
				//all ok, gen user
				u := data.RegistrationUser(login, email, password)
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(u)
				logs.UserCreate(u)
			} else {
				//email bad
				log.Println(customErr.ErrEmailAlreadyUsing.Error())
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, customErr.ErrEmailAlreadyUsing.Error())
			}
		} else {
			log.Println(ue.Login)
			//логин уже есть
			log.Println(customErr.ErrLoginAlreadyUsing)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, customErr.ErrLoginAlreadyUsing.Error())
		}
	} else {
		//bad token
		log.Println(customErr.ErrAccessDenied.Error())
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, customErr.ErrAccessDenied.Error())
	}

}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	//init params
	params := mux.Vars(r)
	criterion := params["criterion"]
	value := params["value"]

	//check rest token
	token := r.Header.Get("acp-token")
	if token == rest.Token {
		user, err := data.FindOne(criterion, value)
		if err != nil {
			//exec error
			log.Println(customErr.ErrNotFound.Error())
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, customErr.ErrNotFound.Error())
		}
		//exec user
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(user)
		logs.UserFound(user)

	} else {
		//return err
		log.Println(customErr.ErrAccessDenied.Error())
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, customErr.ErrAccessDenied.Error())
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//check rest token
	token := r.Header.Get("acp-token")
	if token == rest.Token {
		users := data.FindAll()
		//exec user
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(users)
		logs.UsersFound(users)
	} else {
		//return err
		log.Println(w, customErr.ErrAccessDenied.Error())
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, customErr.ErrAccessDenied.Error())
	}
}

func UcpCreate(w http.ResponseWriter, r *http.Request) {
	//init params
	params := mux.Vars(r)
	login := params["login"]
	nickname := params["nickname"]
	floor := params["floor"]
	age := params["age"]
	nazi := params["nazi"]
	skin := params["skin"]
	country := params["country"]
	quenta := params["quenta"]

	token := r.Header.Get("user-token")
	user, err := data.FindOne("login", login)
	switch err {
	case customErr.ErrNotFound:
		//err
		log.Println(customErr.ErrNotFound.Error())
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, customErr.ErrNotFound.Error())
	case nil:
		//go
		if token == user.AuthToken {
			_, err := data.UcpGetPers(nickname)
			if err == customErr.ErrNotFound {
				_, err := data.UcpCreatePers(nickname, login, floor, age, nazi, skin, country, quenta)
				if err == customErr.ErrUcpFullPerson {
					log.Println(customErr.ErrUcpFullPerson.Error())
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, customErr.ErrUcpFullPerson.Error())
				}
			} else {
				//err person already have
				log.Println(customErr.ErrLoginAlreadyUsing.Error())
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, customErr.ErrLoginAlreadyUsing.Error())
			}
		} else {
			//err bad token
			log.Println(w, customErr.ErrAccessDenied.Error())
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, customErr.ErrAccessDenied.Error())
		}
	default:
		//err
		log.Println(customErr.ErrNotFound.Error())
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, customErr.ErrNotFound.Error())
	}

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
