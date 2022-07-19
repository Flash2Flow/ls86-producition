package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Params struct {
	Url      string
	Database string
	User     string
	Password string
}

type Data struct {
	Params Params
	User   User
	Error  Error
}

type User struct {
	Id          int    `gorm:"primary_key"`
	Login       string `gorm:"not_null"`
	Email       string `gorm:"not_null"`
	Password    string `gorm:"not_null"`
	AuthToken   string `gorm:"not_null"`
	Permissions string `gorm:"not_null"`
	PersOne     string `gorm:"not_null"`
	PersTwo     string `gorm:"not_null"`
	PersThree   string `gorm:"not_null"`
	PersFour    string `gorm:"not_null"`
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

func (d *Data) Connection() (db *gorm.DB) {
	str := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=%s", ls86.Data.Params.User, ls86.Data.Params.Password, ls86.Data.Params.Database, ls86.Data.Params.Url)
	db, err := gorm.Open("mysql",
		str)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("succedssed request...")
	}

	return db
}

func (u *User) Create(use *User) *User {
	db := ls86.Data.Connection()
	user := &User{
		Login:     use.Login,
		Email:     use.Email,
		Password:  use.Password,
		AuthToken: use.AuthToken,
	}
	db.Create(user)
	return user
}

func (u *User) FindAll() []User {
	db := ls86.Data.Connection()

	var users []User
	db.Find(&users)
	fmt.Println(users)
	return users
}

func (u *User) FindOne(title string, value string) (*User, error) {
	switch title {
	case "id":
		db := ls86.Data.Connection()
		user := new(User)
		db.First(user, &User{Login: value})
		fmt.Println(user)
		defer db.Close()
		if user.Id == 0 {
			//err
			return nil, ls86.Data.Error.NotFound
		} else {
			return user, nil
		}

	case "login":
		db := ls86.Data.Connection()
		user := new(User)
		db.First(user, &User{Login: value})
		fmt.Println(user)
		defer db.Close()
		if user.Id == 0 {
			//err
			return nil, ls86.Data.Error.NotFound
		} else {
			return user, nil
		}

	case "email":
		db := ls86.Data.Connection()
		user := new(User)
		db.First(user, &User{Email: value})
		fmt.Println(user)
		defer db.Close()
		if user.Id == 0 {
			//err
			return nil, ls86.Data.Error.NotFound
		} else {
			return user, nil
		}

	case "auth-token":
		db := ls86.Data.Connection()
		user := new(User)
		db.First(user, &User{Login: value})
		fmt.Println(user)
		defer db.Close()
		if user.Id == 0 {
			//err
			return nil, ls86.Data.Error.NotFound
		} else {
			return user, nil
		}
	default:
		return nil, nil
	}
}
