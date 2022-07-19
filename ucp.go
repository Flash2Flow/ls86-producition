package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ucp(w http.ResponseWriter, r *http.Request) {
	uuid, _ := ls86.Cookie.UUID.get(r)
	if uuid != nil {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, have uuid
			id, _ := ls86.Cookie.Id.get(r)
			if id != nil {
				u, _ := ls86.Data.User.FindOne("id", id.Value)
				if u.AuthToken == hash.Value {
					log.Println("USER GET UCP PAGE UUID- " + uuid.Value + " ID- " + id.Value)
					temp, err := template.ParseFiles("temp/ucp.html")

					if err != nil {
						fmt.Fprintf(w, err.Error())
					}

					temp.ExecuteTemplate(w, "ucp", u)
				} else {
					ls86.Redirect.toExit(w)
				}
			} else {
				ls86.Redirect.toExit(w)
			}
		} else {
			//no hash, have uuid
			ls86.Redirect.toExit(w)
		}
	} else {
		hash, _ := ls86.Cookie.Hash.get(r)
		if hash != nil {
			//have hash, no uuid
			log.Println("redirect unknown guest with hash to /exit - " + hash.Value)
			ls86.Redirect.toExit(w)
			return

		} else {
			//no hash no uuid
			ls86.Redirect.toExit(w)
		}
	}
}
