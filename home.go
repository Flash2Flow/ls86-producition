package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func all(w http.ResponseWriter, r *http.Request) {
	/*
		db := ls86.Data.Connection()
		db.CreateTable(User{})
		db.CreateTable(Pers{})
		defer db.Close()
	*/
	bool, u := checkCookie(r)
	if bool == true {
		temp, err := template.ParseFiles("temp/redirects/toHome.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "redirect_home", nil)
		return
	} else {
		temp, err := template.ParseFiles("temp/home.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "home", u)
	}

}

type Lk_Page struct {
	Users User
	P     []Pers
	Count string
	Full  bool
}

func lk(w http.ResponseWriter, r *http.Request) {

	bool, u := checkCookie(r)
	if bool == true {

		//get person
		p, _ := ls86.Data.GetAllPers(u.Login)
		count := GetPersontValue(u)
		if count != "4/4" {
			var full = false
			data := &Lk_Page{Users: *u, P: p, Count: count, Full: full}
			temp, err := template.ParseFiles("temp/lk.html")

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			temp.ExecuteTemplate(w, "lk", data)
		} else {
			var full = true
			data := &Lk_Page{Users: *u, P: p, Count: count, Full: full}
			temp, err := template.ParseFiles("temp/lk.html")

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			temp.ExecuteTemplate(w, "lk", data)
		}

	} else {
		temp, err := template.ParseFiles("temp/redirects/toMain.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "redirect_main", nil)
		return
	}
}

type Lk_Pers_Page struct {
	U User
	P Pers
	S int
}

func lkpers(w http.ResponseWriter, r *http.Request) {

	bool, u := checkCookie(r)
	if bool == true {
		params := mux.Vars(r)
		id := params["id"]
		p, err := ls86.Data.GetPers("id", id)
		if err == nil {
			if p.State == ls86.State.Allow {
				data := &Lk_Pers_Page{U: *u, P: *p, S: 1}

				temp, err := template.ParseFiles("temp/lk_pers.html")

				if err != nil {
					fmt.Fprintf(w, err.Error())
				}

				temp.ExecuteTemplate(w, "lk_pers", data)
			}

			if p.State == ls86.State.Disallow {
				data := &Lk_Pers_Page{U: *u, P: *p, S: 2}

				temp, err := template.ParseFiles("temp/lk_pers.html")

				if err != nil {
					fmt.Fprintf(w, err.Error())
				}

				temp.ExecuteTemplate(w, "lk_pers", data)
			}

			if p.State == ls86.State.Waiting {
				data := &Lk_Pers_Page{U: *u, P: *p, S: 3}

				temp, err := template.ParseFiles("temp/lk_pers.html")

				if err != nil {
					fmt.Fprintf(w, err.Error())
				}

				temp.ExecuteTemplate(w, "lk_pers", data)
			}

		} else {
			temp, err := template.ParseFiles("temp/redirects/toHome.html")

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			temp.ExecuteTemplate(w, "redirect_home", nil)
			return
		}

	} else {
		temp, err := template.ParseFiles("temp/redirects/toMain.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "redirect_main", nil)
		return
	}
}

func ucp(w http.ResponseWriter, r *http.Request) {

	bool, u := checkCookie(r)
	if bool == true {

		temp, err := template.ParseFiles("temp/ucp.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "ucp", u)
	} else {
		temp, err := template.ParseFiles("temp/redirects/toMain.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "redirect_main", nil)
		return
	}
}
