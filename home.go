package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

				} else {
					ls86.Logs.User.Authorized(u)
					temp, err := template.ParseFiles("temp/Home.html")

					if err != nil {
						fmt.Fprintf(w, err.Error())
					}

					temp.ExecuteTemplate(w, "home", u)
				}
			} else {
				//have uuid, have hash, no id
				str := fmt.Sprintf("USER TRYING ACCESS ON LK PAGE WITHOUT ID, but WITH UUID - %s, HASH - %s", uuid.Value, hash.Value)
				log.Println(str)
				ls86.Redirect.toExit(w)
				return
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
