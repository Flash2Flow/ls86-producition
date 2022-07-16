package main

import "errors"

type Error struct {
	PasswordWrong    error
	NotFound         error
	UserTryBadAuth   error
	LoginAlreadyHave error
	EmailAlreadyHave error
	PersAlreadyHave  error
	AccessDenied     error
	FullPersons      error
}

var (
	ls86 = LS86{
		Error: Error{
			PasswordWrong:    errors.New("Password Wrong!"),
			NotFound:         errors.New("Not Found!"),
			UserTryBadAuth:   errors.New("User Try Bad Auth!"),
			LoginAlreadyHave: errors.New("Login Already Have!"),
			EmailAlreadyHave: errors.New("Email Already Have!"),
			AccessDenied:     errors.New("Access Denied!"),
			FullPersons:      errors.New("Full Persons!"),
			PersAlreadyHave:  errors.New("Pers Already Have!"),
		},
		Data: Data{
			User:     "root",
			Password: "537003",
		},
		State: State{
			Allow:    "Allow",
			Disallow: "Disallow",
			Waiting:  "Waiting",
		},
	}
)
