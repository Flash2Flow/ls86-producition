package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Redirect struct {
}

func (r *Redirect) toMain(w http.ResponseWriter) {
	temp, err := template.ParseFiles("temp/redirects/toMain.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "redirect_main", nil)
}

func (r *Redirect) toHome(w http.ResponseWriter) {
	temp, err := template.ParseFiles("temp/redirects/toHome.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "redirect_home", nil)
}

func (r *Redirect) toExit(w http.ResponseWriter) {
	temp, err := template.ParseFiles("temp/redirects/toExit.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "redirect_exit", nil)
}
