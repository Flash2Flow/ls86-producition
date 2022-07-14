package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func ucp(w http.ResponseWriter, r *http.Request) {
	boo, u := checkCookie(r)
	if boo == true {
		temp, err := template.ParseFiles("temp/ucp/three.html", "static/ucp/three/src/app.js", "static/ucp/three/src/jquery-3.6.0.min.js")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		temp.ExecuteTemplate(w, "ucp", u)
	} else {
		//err
		log.Println("LK ERROR")
		temp, err := template.ParseFiles("temp/redirects/toHome.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		temp.ExecuteTemplate(w, "redirect_main", nil)
	}

}
