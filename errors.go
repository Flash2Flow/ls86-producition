package main

import "errors"

type CustomErr struct {
	ErrNotFound          error
	ErrPasswordWrong     error
	ErrAccessDenied      error
	ErrNoErrors          error
	ErrLoginAlreadyUsing error
	ErrEmailAlreadyUsing error
	ErrUcpFullPerson     error
}

var (
	customErr = CustomErr{
		ErrNotFound:          errors.New("Error, not found!"),
		ErrPasswordWrong:     errors.New("Error, password wrong!"),
		ErrAccessDenied:      errors.New("Error, access denied!"),
		ErrLoginAlreadyUsing: errors.New("Error, this login already using!"),
		ErrEmailAlreadyUsing: errors.New("Error, this email already using!"),
		ErrNoErrors:          errors.New("No errors!"),
		ErrUcpFullPerson:     errors.New("Full person!"),
	}
)
