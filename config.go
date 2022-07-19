package main

import "errors"

type LS86 struct {
	Server   Server
	Cookie   Cookie
	Data     Data
	Logs     Logs
	REST     REST
	Redirect Redirect
}

type Server struct {
	Test  bool
	Url   string
	Close bool
	Port  string
}

var (
	ls86 = LS86{
		Server: Server{
			Test:  true,
			Close: false,
			Url:   "localhost",
			Port:  ":8080",
		},
		Data: Data{
			Params: Params{
				Url:      "Local",
				Database: "ls86",
				User:     "root",
				Password: "537003Sa$",
			},
			Error: Error{
				NotFound:         errors.New("Not Found!"),
				TokenBad:         errors.New("Token Bad!"),
				LoginAlreadyHave: errors.New("Email Already Have!"),
				EmailAlreadyHave: errors.New("Login Already Have!"),
			}},
		Redirect: Redirect{},
	}
)
