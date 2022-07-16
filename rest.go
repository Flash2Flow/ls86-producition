package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UcpCreate(w http.ResponseWriter, r *http.Request) {
	//init values
	params := mux.Vars(r)
	id := params["id"]
	name := params["name"]
	floor := params["floor"]
	age := params["age"]
	nazi := params["nazi"]
	skin := params["skin"]
	country := params["country"]
	quenta := params["quenta"]
	token := r.Header.Get("user-token")

	u, err := ls86.Data.GetUser("id", id)
	switch err {
	case ls86.Error.NotFound:
		log.Println(ls86.Error.NotFound)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, ls86.Error.NotFound.Error())
	case nil:
		if u.AuthToken == token {
			_, err := ls86.Data.GetPers("name", name)
			if err == nil {
				log.Println(ls86.Error.PersAlreadyHave)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, ls86.Error.PersAlreadyHave.Error())
			} else {
				val := GetPersontValue(u)
				switch val {
				case "4/4":
					//err
					log.Println(ls86.Error.FullPersons)
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, ls86.Error.FullPersons.Error())
				case "3/4":
					person := Pers{
						Name:     name,
						Login:    u.Login,
						Floor:    floor,
						Age:      age,
						Nazi:     nazi,
						Skin:     skin,
						Country:  country,
						Quenta:   quenta,
						State:    ls86.State.Waiting,
						State_Id: "4",
					}
					ls86.Data.UserUpdate("pers_four", u.Login, name)
					p := ls86.Data.CreatePers(&person)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(p)

				case "2/4":
					person := Pers{
						Name:     name,
						Login:    u.Login,
						Floor:    floor,
						Age:      age,
						Nazi:     nazi,
						Skin:     skin,
						Country:  country,
						Quenta:   quenta,
						State:    ls86.State.Waiting,
						State_Id: "3",
					}
					ls86.Data.UserUpdate("pers_three", u.Login, name)
					p := ls86.Data.CreatePers(&person)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(p)

				case "1/4":
					person := Pers{
						Name:     name,
						Login:    u.Login,
						Floor:    floor,
						Age:      age,
						Nazi:     nazi,
						Skin:     skin,
						Country:  country,
						Quenta:   quenta,
						State:    ls86.State.Waiting,
						State_Id: "2",
					}
					ls86.Data.UserUpdate("pers_two", u.Login, name)
					p := ls86.Data.CreatePers(&person)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(p)

				case "0/4":
					person := Pers{
						Name:     name,
						Login:    u.Login,
						Floor:    floor,
						Age:      age,
						Nazi:     nazi,
						Skin:     skin,
						Country:  country,
						Quenta:   quenta,
						State:    ls86.State.Waiting,
						State_Id: "1",
					}
					ls86.Data.UserUpdate("pers_one", u.Login, name)
					p := ls86.Data.CreatePers(&person)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(p)
				}
			}

		} else {
			//err
			log.Println(ls86.Error.AccessDenied)
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, ls86.Error.AccessDenied.Error())
		}
	default:
		log.Println(err)
	}
}

func authAjax(w http.ResponseWriter, r *http.Request) {
	//init block
	params := mux.Vars(r)
	login := params["login"]
	password := params["password"]

	u, err := ls86.Data.Auth(login, password)
	if err == nil {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(u)
	} else {
		log.Println(ls86.Error.PasswordWrong.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, ls86.Error.PasswordWrong.Error())
	}

}
