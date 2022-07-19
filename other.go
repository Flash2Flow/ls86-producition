package main

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetPersontValue(u *User) string {
	if u.PersOne == "" {
		return "0/4"
	} else {
		if u.PersTwo == "" {
			return "1/4"
		} else {
			if u.PersThree == "" {
				return "2/4"
			} else {
				if u.PersFour == "" {
					return "3/4"
				} else {
					return "4/4"
				}
			}
		}
	}
}
