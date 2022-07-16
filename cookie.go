package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func auth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	login := params["login"]
	authToken := params["token"]

	if login != " " {
		if authToken != " " {
			u, err := ls86.Data.GetUser("login", login)
			if err == nil {
				if authToken == u.AuthToken {
					str := strconv.Itoa(u.Id)
					Id := &http.Cookie{
						Name:     "Id",
						Value:    str,
						MaxAge:   3000,
						Domain:   "ls-86-rp.ru",
						SameSite: http.SameSiteLaxMode,
						Path:     "/",
					}
					hash := &http.Cookie{
						Name:     "hash",
						Value:    u.AuthToken,
						Domain:   "ls-86-rp.ru",
						Path:     "/",
						SameSite: http.SameSiteLaxMode,
						MaxAge:   3000,
					}

					http.SetCookie(w, Id)
					http.SetCookie(w, hash)
					temp, err := template.ParseFiles("temp/redirects/toHome.html")

					if err != nil {
						fmt.Fprintf(w, err.Error())
					}

					temp.ExecuteTemplate(w, "redirect_home", nil)
					return
				} else {
					log.Println(ls86.Error.UserTryBadAuth)
					log.Println("user- " + u.AuthToken)
					log.Println("token- " + authToken)
				}

			}
		}
	}
}

func reg(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	login := params["login"]
	email := params["email"]
	password := params["password"]
	_, err := ls86.Data.GetUser("login", login)
	if err != ls86.Error.LoginAlreadyHave {
		_, err := ls86.Data.GetUser("login", email)
		if err != ls86.Error.EmailAlreadyHave {
			r := ls86.Data.Registration(login, email, password)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(r)
		} else {
			//email уже есть
			log.Println(err)
		}
	} else {
		//login уже есть
		log.Println(err)
	}
}

func exit(w http.ResponseWriter, r *http.Request) {
	Id := &http.Cookie{
		Name:     "Id",
		Value:    "",
		MaxAge:   -1,
		Domain:   "ls-86-rp.ru",
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	hash := &http.Cookie{
		Name:     "hash",
		Value:    "",
		Domain:   "ls-86-rp.ru",
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}

	http.SetCookie(w, Id)
	http.SetCookie(w, hash)

	temp, err := template.ParseFiles("temp/redirects/toMain.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "redirect_main", nil)
	return
}

func checkCookie(r *http.Request) (bool, *User) {
	tokenCookie, err := r.Cookie("hash")
	if err == nil {
		id, err2 := r.Cookie("Id")
		if err2 == nil {
			if tokenCookie.Value != " " {
				if id.Value != " " {
					u, err := ls86.Data.GetUser("id", id.Value)
					switch err {
					case ls86.Error.NotFound:
						//err not found
						log.Println(ls86.Error.NotFound)
						return false, nil

					case nil:
						if tokenCookie.Value == u.AuthToken {
							return true, u
						} else {
							//bad token
							log.Println("bad token")
							return false, nil
						}
					default:
						//err unknown
						log.Println(err)
						return false, nil
					}

				} else {
					//empty id
					log.Println("empty id")
					return false, nil
				}
			} else {
				//empty token
				log.Println("empty token")
				return false, nil
			}
		} else {
			//err cookie id
			log.Println("err cookie id")
			return false, nil
		}
	} else {
		//err cookie token
		log.Println("err cookie token")
		return false, nil
	}

}
