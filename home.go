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
						Value:    u.AuthToken,
						Domain:   "localhost",
						Path:     "/",
						SameSite: http.SameSiteLaxMode,
						MaxAge:   300,
					}

					http.SetCookie(w, Id)
					http.SetCookie(w, hash)
					temp, err := template.ParseFiles("temp/test.html")

					if err != nil {
						fmt.Fprintf(w, err.Error())
					}

					temp.ExecuteTemplate(w, "redirect", nil)
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

func lkuser(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("hash")
	if err == nil {
		id, err2 := r.Cookie("ID")
		if err2 == nil {
			if tokenCookie.Value != " " {
				if id.Value != " " {
					u, err := data.FindOne("id", id.Value)
					switch err {
					case customErr.ErrNotFound:
						//err not found
						log.Println(customErr.ErrNotFound)

					case nil:
						if tokenCookie.Value == u.AuthToken {

							temp, err := template.ParseFiles("temp/lk_user.html", "static/lk_user/src/app.js", "static/lk_user/src/jquery-3.6.0.min.js")

							if err != nil {
								fmt.Fprintf(w, err.Error())
							}

							temp.ExecuteTemplate(w, "lk_user", u)
						} else {
							//bad token
							log.Println("bad token")
						}
					default:
						//err unknown
						log.Println(err)
					}

				} else {
					//empty id
					log.Println("empty id")
				}
			} else {
				//empty token
				log.Println("empty token")
			}
		} else {
			//err cookie id
			log.Println("err cookie id")
		}
	} else {
		//err cookie token
		log.Println("err cookie token")
	}

}

func lk(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("hash")
	if err == nil {
		id, err2 := r.Cookie("ID")
		if err2 == nil {
			if tokenCookie.Value != " " {
				if id.Value != " " {
					u, err := data.FindOne("id", id.Value)
					switch err {
					case customErr.ErrNotFound:
						//err not found
						log.Println(customErr.ErrNotFound)

					case nil:
						if tokenCookie.Value == u.AuthToken {

							temp, err := template.ParseFiles("temp/lk.html", "static/lk/src/app.js", "static/lk/src/jquery-3.6.0.min.js")

							if err != nil {
								fmt.Fprintf(w, err.Error())
							}

							temp.ExecuteTemplate(w, "lk", u)
						} else {
							//bad token
							log.Println("bad token")
						}
					default:
						//err unknown
						log.Println(err)
					}

				} else {
					//empty id
					log.Println("empty id")
				}
			} else {
				//empty token
				log.Println("empty token")
			}
		} else {
			//err cookie id
			log.Println("err cookie id")
		}
	} else {
		//err cookie token
		log.Println("err cookie token")
	}

}
