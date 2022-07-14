package main

import (
	"fmt"
	"log"
	"time"
)

type Logs struct {
	Server Server
	Api    Api
}

type Api struct {
	FoundUser    string
	FoundUsers   string
	CreateUser   string
	CreatePerson string
	UpdatePerson string
	DeletePerson string
}

type Server struct {
	Starting string
	Port     string
	Dev      Dev
	Version  string
}

type Dev struct {
	Company  string
	Backend  string
	Frontend string
	URL      string
}

var (
	logs = Logs{
		Server: Server{
			Port:     ":3030",
			Starting: "server starting...",
			Version:  "0.1 beta",
			Dev: Dev{
				Company:  "vavylon org",
				Backend:  ".re-incarnation",
				Frontend: ".shot",
				URL:      "https://vavylon.com",
			},
		},
		Api: Api{
			FoundUser:    "User has been found: " + time.Now().Format("2006-01-02 15:04:05") + "\r\n",
			FoundUsers:   "Users has been found, row list: " + time.Now().Format("2006-01-02 15:04:05") + "\r\n",
			CreateUser:   "User has been created: " + time.Now().Format("2006-01-02 15:04:05") + "\r\n",
			CreatePerson: "Person has been created: " + time.Now().Format("2006-01-02 15:04:05") + "\r\n",
			UpdatePerson: "Person has been update: " + time.Now().Format("2006-01-02 15:04:05") + "\r\n",
			DeletePerson: "Person has been delete: " + time.Now().Format("2006-01-02 15:04:05") + "\r\n",
		},
	}
)

func (l Logs) Logging(value string) {
	switch value {
	case logs.Server.Starting:
		text := fmt.Sprintf("%s\r\nPort: %s\r\ndev by: %s, %s \r\n%s\r\n%s\r\n%s", logs.Server.Starting, logs.Server.Port, logs.Server.Dev.Backend, logs.Server.Dev.Frontend, logs.Server.Dev.Company, logs.Server.Dev.URL, logs.Server.Version)
		log.Println(text)
	}
}

func (l Logs) UserFound(u *User) {
	text := fmt.Sprintf("%s%v", logs.Api.FoundUser, u)
	log.Println(text)
}

func (l Logs) UsersFound(u []User) {
	text := fmt.Sprintf("%s%v", logs.Api.FoundUsers, u)
	log.Println(text)
}

func (l Logs) UserCreate(u *User) {
	text := fmt.Sprintf("%s%v", logs.Api.CreateUser, u)
	log.Println(text)
}

func (l Logs) CreatePerson(p *Person) {
	text := fmt.Sprintf("%s%v", logs.Api.CreatePerson, p)
	log.Println(text)
}

func (l Logs) UpdatePerson(p *Person) {
	text := fmt.Sprintf("%s%v", logs.Api.UpdatePerson, p)
	log.Println(text)
}

func (l Logs) DeletePerson(p *Person) {
	text := fmt.Sprintf("%s%v", logs.Api.DeletePerson, p)
	log.Println(text)
}
