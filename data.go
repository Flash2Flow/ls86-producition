package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Data struct {
	User     string
	Password string
}

type User struct {
	Id        int    `gorm:"primary_key"`
	Login     string `gorm:"not_null"`
	Email     string `gorm:"not_null"`
	Password  string `gorm:"not_null"`
	AuthToken string `gorm:"not_null"`
}

var (
	data = Data{
		User:     "root",
		Password: "537003",
	}
)

func (d *Data) Connection() (db *gorm.DB) {
	db, err := gorm.Open("mysql",
		"root:537003@/ls86?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("succedssed request...")
	}

	//defer db.Close()
	return db
}

func (d *Data) FindOne(title string, value string) (*User, error) {

	switch title {
	case "login":
		db := data.Connection()
		user := new(User)
		db.First(user, &User{Login: value})
		if user.Id == 0 {
			defer db.Close()
			return nil, customErr.ErrNotFound
		} else {
			defer db.Close()
			return user, nil
		}

	case "email":
		db := data.Connection()
		user := new(User)
		db.First(user, &User{Email: value})
		if user.Id == 0 {
			defer db.Close()
			return nil, customErr.ErrNotFound
		} else {
			defer db.Close()
			return user, nil
		}

	case "id":
		db := data.Connection()
		user := new(User)
		intvalue, err := strconv.Atoi(value)
		if err != nil {
			log.Println("Idiot err")
		}
		db.First(user, &User{Id: intvalue})
		if user.Id == 0 {
			defer db.Close()
			return nil, customErr.ErrNotFound
		} else {
			defer db.Close()
			return user, nil
		}
	}

	return nil, customErr.ErrNotFound
}

func (d *Data) FindAll() []User {
	db := data.Connection()
	var users []User
	db.Find(&users)

	defer db.Close()
	return users
}

func (d *Data) RegistrationUser(login string, email string, password string) *User {
	//gen password
	hash, _ := HashPassword(password)
	db := data.Connection()
	rand := RandStringRunes(32)
	//db.CreateTable(User{})
	user := &User{Login: login, Email: email, Password: hash, AuthToken: rand}
	db.Create(user)

	defer db.Close()
	return user
}

func (d *Data) Auth(login string, password string) (*User, error) {
	log.Println(login)
	u, err := data.FindOne("login", login)
	if err != nil {
		return nil, customErr.ErrNotFound
	} else {
		passBool := CheckPasswordHash(password, u.Password)

		if passBool == true {
			return u, nil
		} else {
			return nil, customErr.ErrPasswordWrong
		}
	}

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
