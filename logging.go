package main

import (
	"fmt"
	"log"
)

type Logs struct {
	User   User
	Data   Data
	Server Server
}

func (u *User) StartServer() {
	log.Println("STARTING SERVER...")
}

func (u *User) GenerateUUID(value string) {
	str := fmt.Sprintf("GENERATE NEW UUID: %s", value)
	log.Println(str)
}

func (u *User) MoveUser(value string, where string, to string) {
	str := fmt.Sprintf("MOVE USER: %s,\r\nFROM: %s -> TO: %s", value, where, to)
	log.Println(str)
}

func (u *User) MoveUUID(value string, where string, to string) {
	str := fmt.Sprintf("MOVE UUID: %s,\r\nFROM: %s -> TO: %s", value, where, to)
	log.Println(str)
}

func (u *User) UserGo(value string, where string) {
	str := fmt.Sprintf("USER GO: %s,\r\nON PAGE: %s", value, where)
	log.Println(str)
}

func (u *User) UUIDGo(value string, where string) {
	str := fmt.Sprintf("USER GO: %s,\r\nON PAGE: %s", value, where)
	log.Println(str)
}

func (u *User) Unauthorized(value string, uuid string) {
	str := fmt.Sprintf("USER LEFT: %s,\r\nUUID: %s", value, uuid)
	log.Println(str)
}

func (u *User) Auth(user *User) {
	str := fmt.Sprintf("REST API AUTH USER: %v", user)
	log.Println(str)
}

func (s *Server) AllOk() {
	log.Println("ALL SERVICES OK, NEXT STEP --->")
}

func (e *Error) NotFoundReq(value string) {
	str := fmt.Sprintf("REST API NOT FOUND REQUEST: %s", value)
	log.Println(str)
}

func (e *Error) IncorectToken(value string) {
	str := fmt.Sprintf("REST API AUTH BAD REQUEST TOKEN: %s", value)
	log.Println(str)
}

func (e *Error) StatusBadRequest(value string) {
	str := fmt.Sprintf("REST API BAD REQUEST: %s", value)
	log.Println(str)
}

func (e *Error) UserLoginAlreadyUsed(value string) {
	str := fmt.Sprintf("REST API BAD REQUEST, LOGIN USED: %s", value)
	log.Println(str)
}

func (e *Error) UserEmailAlreadyUsed(value string) {
	str := fmt.Sprintf("REST API BAD REQUEST, EMAIL USED: %s", value)
	log.Println(str)
}

func (e *Error) PasswordBad(value string) {
	str := fmt.Sprintf("REST API BAD REQUEST, PASSWORD INCORECT: %s", value)
	log.Println(str)
}

func (s *Success) Found(value interface{}) {
	str := fmt.Sprintf("REST API FOUND USER: %v", value)
	log.Println(str)
}

func (s *Success) Created(value interface{}) {
	str := fmt.Sprintf("REST API CREATE USER: %v", value)
	log.Println(str)
}

func (s *Success) Authorized(value interface{}) {
	str := fmt.Sprintf("REST API AUTH USER: %v", value)
	log.Println(str)
}
