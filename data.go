package main

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LS86 struct {
	Error Error
	Data  Data
	State State
}

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
	PersOne   string `gorm:"not_null"`
	PersTwo   string `gorm:"not_null"`
	PersThree string `gorm:"not_null"`
	PersFour  string `gorm:"not_null"`
}

type Pers struct {
	Id       int    `gorm:"primary_key"`
	Name     string `gorm:"not_null"`
	Login    string `gorm:"not_null"`
	Floor    string `gorm:"not_null"`
	Age      string `gorm:"not_null"`
	Nazi     string `gorm:"not_null"`
	Skin     string `gorm:"not_null"`
	Country  string `gorm:"not_null"`
	Quenta   string `gorm:"not_null"`
	State    string `gorm:"not_null"`
	State_Id string `gorm:"not_null"`
}

type State struct {
	Allow    string
	Disallow string
	Waiting  string
}

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

func (d *Data) Registration(login string, email string, password string) *User {
	//gen password
	hash, _ := HashPassword(password)
	db := ls86.Data.Connection()
	rand := RandStringRunes(32)
	//db.CreateTable(User{})
	user := &User{Login: login, Email: email, Password: hash, AuthToken: rand}
	db.Create(user)

	defer db.Close()
	return user
}

func (d *Data) Auth(login string, password string) (*User, error) {
	u, err := ls86.Data.GetUser("login", login)
	if err != nil {
		return nil, err
	} else {
		passBool := CheckPasswordHash(password, u.Password)

		if passBool == true {
			return u, nil
		} else {
			return nil, ls86.Error.PasswordWrong
		}
	}

}

func (d *Data) GetUser(title string, value string) (*User, error) {
	switch title {
	case "login":
		db := ls86.Data.Connection()
		user := new(User)
		db.First(user, &User{Login: value})
		defer db.Close()
		if user.Id == 0 {
			return nil, ls86.Error.NotFound
		} else {
			return user, nil
		}

	case "email":
		db := ls86.Data.Connection()
		user := new(User)
		db.First(user, &User{Email: value})
		defer db.Close()
		if user.Id == 0 {
			return nil, ls86.Error.NotFound
		} else {
			return user, nil
		}
	case "id":
		db := ls86.Data.Connection()
		user := new(User)
		str, _ := strconv.Atoi(value)
		db.First(user, &User{Id: str})
		defer db.Close()
		if user.Id == 0 {
			return nil, ls86.Error.NotFound
		} else {
			return user, nil
		}
	default:
		return nil, nil
	}
}

//UserUpdate

func (d *Data) UserUpdate(title string, login string, value string) {
	switch title {
	case "pers_one":
		db := ls86.Data.Connection()
		user := &User{Login: login}
		db.Model(user).Update("pers_one", value)
		defer db.Close()
	case "pers_two":
		db := ls86.Data.Connection()
		user := &User{Login: login}
		db.Model(user).Update("pers_two", value)
		defer db.Close()
	case "pers_three":
		db := ls86.Data.Connection()
		user := &User{Login: login}
		db.Model(user).Update("pers_three", value)
		defer db.Close()
	case "pers_four":
		db := ls86.Data.Connection()
		user := &User{Login: login}
		db.Model(user).Update("pers_four", value)
		defer db.Close()
	}

}

func (d *Data) CreatePers(p *Pers) *Pers {
	db := ls86.Data.Connection()
	pers := &p
	db.Create(pers)

	defer db.Close()
	return p
}

func (d *Data) GetPers(title string, value string) (*Pers, error) {
	switch title {

	case "id":
		db := ls86.Data.Connection()
		pers := new(Pers)
		str, _ := strconv.Atoi(value)
		db.First(pers, &Pers{Id: str})
		defer db.Close()
		if pers.Id == 0 {
			return nil, ls86.Error.NotFound
		} else {
			return pers, nil
		}

	case "name":
		db := ls86.Data.Connection()
		pers := new(Pers)
		db.First(pers, &Pers{Name: value})
		defer db.Close()
		if pers.Id == 0 {
			return nil, ls86.Error.NotFound
		} else {
			return pers, nil
		}

	default:
		return nil, nil

	}
}

func (d *Data) GetAllPers(value string) ([]Pers, error) {
	db := ls86.Data.Connection()

	var person []Pers
	db.Find(&person)
	defer db.Close()
	for _, p := range person {
		if p.Login == value {
			//fmt.Println(person)
			return person, nil
		}
	}
	return nil, ls86.Error.NotFound
}
