package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func all(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("temp/home.html", "static/src/app.js", "static/src/jquery-3.6.0.min.js")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "home", nil)

}

func auth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	login := params["login"]
	authToken := params["token"]

	if login != " " {
		if authToken != " " {
			u, err := data.FindOne("login", login)
			if err == nil {
				if authToken == u.AuthToken {
					str := strconv.Itoa(u.Id)
					Id := &http.Cookie{
						Name:     "ID",
						Value:    str,
						MaxAge:   300,
						Domain:   "localhost",
						SameSite: http.SameSiteLaxMode,
						Path:     "/",
					}
					hash := &http.Cookie{
						Name:     "hash",
						Value:    u.Password,
						Domain:   "localhost",
						Path:     "/",
						SameSite: http.SameSiteLaxMode,
						MaxAge:   300,
					}

					http.SetCookie(w, Id)
					http.SetCookie(w, hash)
					w.WriteHeader(200)
					w.Write([]byte("Doc Get Successful"))
					return
				} else {
					log.Println(1)
					log.Println("u- " + u.AuthToken)
					log.Println("p- " + authToken)
				}

			}
		}
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Cookie("ID")

	u, err := data.FindOne("id", id.Value)
	if err != nil {
		//err
		log.Println(err)
	}
	//user := fmt.Sprintf("%v", u)
	fmt.Fprintf(w, u.Login)

	/*
		tokenCookie, err := r.Cookie("hash")
		id, err2 := r.Cookie("ID")

		if err != nil {
			if tokenCookie.Value != " " {
				if err2 != nil {
					if id.Value != " " {
						u, err := data.FindOne("id", id.Value)
						if err != nil {
							//err
							log.Println(err)
						}
						if tokenCookie.Value == u.AuthToken {
							//okey
							user := fmt.Sprintf("%v", u)
							fmt.Fprintf(w, user)
						} else {
							//err bad token
							log.Println("bad token")
						}
					} else {
						//err empty id
						log.Println("empty id")
					}
				} else {
					//err cookie id
					log.Println("err cookie id")
				}
			} else {
				//err empty token
				log.Println("err empty token")
			}
		} else {
			//err cookie token
			log.Println(err2)
		}
	*/
}
