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
	Id          int    `gorm:"primary_key"`
	Login       string `gorm:"not_null"`
	Email       string `gorm:"not_null"`
	Password    string `gorm:"not_null"`
	AuthToken   string `gorm:"not_null"`
	PersonOne   string `gorm:"not_null"`
	PersonTwo   string `gorm:"not_null"`
	PersonThree string `gorm:"not_null"`
	PersonFour  string `gorm:"not_null"`
}

type Person struct {
	Id      int    `gorm:"primary_key"`
	Name    string `gorm:"not_null"`
	Login   string `gorm:"not_null"`
	Floor   string `gorm:"not_null"`
	Age     string `gorm:"not_null"`
	Nazi    string `gorm:"not_null"`
	Skin    string `gorm:"not_null"`
	Country string `gorm:"not_null"`
	Quenta  string `gorm:"not_null"`
	State   string `gorm:"not_null"`
}

type State struct {
	Waiting  string
	Allow    string
	Disallow string
}

var (
	data = Data{
		User:     "root",
		Password: "537003",
	}
	state = State{
		Waiting:  "Waiting",
		Allow:    "Allow",
		Disallow: "Disallow",
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

	case "token":
		db := data.Connection()
		user := new(User)
		db.First(user, &User{AuthToken: value})
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

func (d *Data) UcpGetPers(value string) (*Person, error) {
	db := data.Connection()
	person := new(Person)
	db.First(person, &Person{Name: value})
	if person.Id == 0 {
		defer db.Close()
		return nil, customErr.ErrNotFound
	} else {
		defer db.Close()
		return person, nil
	}

}

func getEmptyPers(u *User) string {
	if u.PersonOne == "" {
		return "PersonOne"
	} else {
		if u.PersonTwo == "" {
			return "PersonTwo"
		} else {
			if u.PersonThree == "" {
				return "PersonThree"
			} else {
				if u.PersonFour == "" {
					return "PersonFour"
				} else {
					return "PersonFull"
				}
			}
		}
	}
}

func (d *Data) UpdateUser(title string, login string, person string, value string) error {
	switch title {
	case "UserUcp":
		switch person {
		case "PersonOne":
			_, err := data.FindOne("login", login)
			if err != customErr.ErrNotFound {
				db := data.Connection()
				user := &User{Login: login}
				db.Model(user).Update("PersonOne", value)
			} else {
				log.Println(err)
				return err
			}
		case "PersonTwo":
			_, err := data.FindOne("login", login)
			if err != customErr.ErrNotFound {
				db := data.Connection()
				user := &User{Login: login}
				db.Model(user).Update("PersonTwo", value)
			} else {
				log.Println(err)
				return err
			}
		case "PersonThree":
			_, err := data.FindOne("login", login)
			if err != customErr.ErrNotFound {
				db := data.Connection()
				user := &User{Login: login}
				db.Model(user).Update("PersonThree", value)
			} else {
				log.Println(err)
				return err
			}
		case "PersonFour":
			_, err := data.FindOne("login", login)
			if err != customErr.ErrNotFound {
				db := data.Connection()
				user := &User{Login: login}
				db.Model(user).Update("PersonFour", value)
			} else {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func (d *Data) UcpLimitPers(login string) (string, error) {
	u, err := data.FindOne("login", login)
	switch err {
	case customErr.ErrNotFound:
		return " ", customErr.ErrNotFound
	case nil:
		str := getEmptyPers(u)
		return str, nil
	default:
		log.Println(err)
	}

	return " ", nil
}

func (d *Data) UcpCreatePers(nickname string, login string, floor string, age string, nazi string, skin string, country string, quenta string) (*Person, error) {
	db := data.Connection()
	//db.CreateTable(Person{})
	persVal, _ := data.UcpLimitPers(login)
	if persVal == "PersonFull" {
		return nil, customErr.ErrUcpFullPerson
	} else {
		data.UpdateUser("UserUcp", login, persVal, nickname)

		person := &Person{Name: nickname, Login: login, Floor: floor, Age: age, Nazi: nazi, Skin: skin, Country: country, Quenta: quenta, State: state.Waiting}
		db.Create(person)

		defer db.Close()
		logs.CreatePerson(person)
		return person, nil
	}
}

func (d *Data) UcpUpdatePers(login string, state string) {
	db := data.Connection()
	person := &Person{Login: login}
	db.Model(person).Update("State", state)
	logs.UpdatePerson(person)

}

func (d *Data) UcpDeletePers(login string) {
	db := data.Connection()
	person := &Person{Login: login}
	db.Delete(person)
	logs.DeletePerson(person)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
