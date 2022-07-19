package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func all(w http.ResponseWriter, r *http.Request) {
	uuid, _ := ls86.Cookie.UUID.get(r)
	if uuid != nil {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, have uuid
			ls86.Logs.User.MoveUUID(uuid.Value, "/", "/home")
			ls86.Redirect.toHome(w)
			return
		} else {
			//no hash, have uuid
			temp, err := template.ParseFiles("temp/home.html")

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			temp.ExecuteTemplate(w, "home", nil)
		}
	} else {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, no uuid
			log.Println("redirect unknown guest with hash to /home - " + hash.Value)
			ls86.Redirect.toExit(w)
			return

		} else {
			//no hash no uuid
			ls86.Logs.User.GenerateUUID(ls86.Cookie.UUID.set(w))
			temp, err := template.ParseFiles("temp/home.html")

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			temp.ExecuteTemplate(w, "home", nil)
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	uuid, _ := ls86.Cookie.UUID.get(r)
	if uuid != nil {
		//have uuid
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have uuid + hash
			id, _ := ls86.Cookie.Id.get(r)
			if id != nil {
				//all ok
				//have uuid + hash + id
				u, err := ls86.Cookie.Other.CheckAuth(id, hash)
				if err != nil {
					log.Println("eererer")
				} else {
					ls86.Logs.User.Auth(u)
					p, _ := ls86.Data.Pers.GetAll(u.Login)
					count := GetPersontValue(u)
					if count != "4/4" {
						data := LKPAGE{
							Users: *u,
							P:     p,
							Count: count,
							Full:  false,
						}
						temp, err := template.ParseFiles("temp/lk.html")

						if err != nil {
							fmt.Fprintf(w, err.Error())
						}

						temp.ExecuteTemplate(w, "lk", data)
					} else {
						data := LKPAGE{
							Users: *u,
							P:     p,
							Count: count,
							Full:  false,
						}
						temp, err := template.ParseFiles("temp/lk.html")

						if err != nil {
							fmt.Fprintf(w, err.Error())
						}

						temp.ExecuteTemplate(w, "lk", data)
					}
				}
			} else {
				//have uuid, have hash, no id
				id, err := ls86.Data.User.FindOne("token", hash.Value)
				if err != nil {
					str := fmt.Sprintf("USER TRYING ACCESS ON LK PAGE WITHOUT ID, but WITH UUID - %s, HASH - %s", uuid.Value, hash.Value)
					log.Println(str)
					ls86.Redirect.toExit(w)
					return
				} else {
					str := fmt.Sprintf("%v", id.Id)
					ls86.Cookie.Id.set(w, str)
				}
			}
		} else {
			//no hash
			id, _ := ls86.Cookie.Id.get(r)
			if id != nil {
				//have uuid, no hash, have id
				str := fmt.Sprintf("USER TRYING ACCESS ON LK PAGE WITHOUT UUID, but WITH UUID - %s, ID - %s", uuid.Value, id.Value)
				log.Println(str)
				ls86.Redirect.toExit(w)
				return
			} else {
				//have uuid, no hash, no id
				log.Println("USER TRYING ACCESS ON LK PAGE WITHOUT HASH, ID but HAVE UUID - " + uuid.Value)
				ls86.Redirect.toMain(w)
				return

			}
		}
	} else {
		//no uuid
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//no uuid, have hash
			id, _ := ls86.Cookie.Id.get(r)
			if id != nil {
				str := fmt.Sprintf("USER TRYING ACCESS ON LK PAGE WITHOUT UUID, but WITH HASH - %s, ID - %s", hash.Value, id.Value)
				//no uuid, have hash, have id
				log.Println(str)
				ls86.Redirect.toExit(w)
				return
			} else {
				//no uuid, have hash, no id
				log.Println("USER TRYING ACCESS ON LK PAGE WITHOUT UUID, ID  but WITH HASH - " + hash.Value)
				ls86.Redirect.toExit(w)
				return
			}
		} else {
			//no uuid, no hash
			id, _ := ls86.Cookie.Id.get(r)
			if id != nil {
				//no uuid, no hash, have id
				log.Println("GUEST TRYING ACCESS ON LK PAGE WITHOUT UUID, HASH but WITH ID - " + id.Value)
				ls86.Redirect.toExit(w)
				return
			} else {
				//no uuid, no hash, no id
				log.Println("GUEST TRYING ACCESS ON LK PAGE WITHOUT UUID, HASH, ID")
				ls86.Redirect.toMain(w)
				return
			}
		}

	}
}

func auth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//url params
	login := params["login"]
	token := params["token"]
	uuid, _ := ls86.Cookie.UUID.get(r)
	if uuid != nil {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, have uuid
			ls86.Logs.User.MoveUUID(uuid.Value, "/", "/home")
			ls86.Redirect.toHome(w)
			return
		} else {
			//no hash, have uuid
			u, err := ls86.Data.User.FindOne("login", login)
			if err != nil {
				log.Println("NOT FOUND USER WITH LOGIN - " + login)
			} else {
				if login == u.Login {
					if token == u.AuthToken {
						str := fmt.Sprintf("%v", u.Id)
						ls86.Cookie.Id.set(w, str)
						ls86.Cookie.Hash.set(w, token)
						ls86.Redirect.toHome(w)
					} else {
						ls86.Data.Error.IncorectToken(token)
					}
				} else {
					log.Println("some errors 👉🏻👈🏻")
				}
			}
		}
	} else {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, no uuid
			log.Println("redirect unknown guest with hash to /home - " + hash.Value)
			ls86.Redirect.toExit(w)
			return

		} else {
			//no hash no uuid
			ls86.Redirect.toMain(w)
		}
	}
}

func exit(w http.ResponseWriter, r *http.Request) {
	uuid, _ := ls86.Cookie.UUID.get(r)
	if uuid != nil {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, have uuid
			ls86.Logs.User.Unauthorized(hash.Value, uuid.Value)

			//del all cookie

			ls86.Cookie.Hash.delete(w)
			ls86.Cookie.Id.delete(w)

			ls86.Redirect.toMain(w)
			return
		} else {
			//no hash, have uuid
			ls86.Logs.User.MoveUUID(uuid.Value, "/exit", "/")
			//termitate id cookie
			ls86.Cookie.Id.delete(w)
			ls86.Cookie.Hash.delete(w)

			ls86.Redirect.toMain(w)
			return
		}
	} else {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, no uuid
			//termitate hash, then redirect to home
			ls86.Cookie.Id.delete(w)
			ls86.Cookie.Hash.delete(w)
			ls86.Redirect.toMain(w)
			return

		} else {
			//no hash no uuid
			log.Println("unknown guest without cookies trying exit...")
			ls86.Cookie.Id.delete(w)
			ls86.Redirect.toMain(w)
			return
		}
	}
}
