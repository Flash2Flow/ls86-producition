package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Cookie struct {
	UUID  UUID
	Id    Id
	Hash  Hash
	Other Other
}

type UUID struct {
}
type Id struct {
}
type Hash struct {
}
type Other struct {
}

func (u *UUID) set(w http.ResponseWriter) string {
	str := uuid.New()
	log.Println(str.String())
	uuid := http.Cookie{
		Name:     "uuid",
		Value:    str.String(),
		Domain:   ls86.Server.Url,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3000,
	}
	http.SetCookie(w, &uuid)
	return str.String()
}

func (u *UUID) get(r *http.Request) (*http.Cookie, error) {
	uuid, err := r.Cookie("uuid")
	if err == nil {
		return uuid, err
	} else {
		return nil, err
	}

}

func (u *UUID) delete(w http.ResponseWriter) {
	UUID := &http.Cookie{
		Name:     "uuid",
		Value:    "",
		MaxAge:   -1,
		Domain:   ls86.Server.Url,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	http.SetCookie(w, UUID)
}

func (i *Id) set(w http.ResponseWriter, value string) {
	Id := http.Cookie{
		Name:     "id",
		Value:    value,
		Domain:   ls86.Server.Url,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3000,
	}
	http.SetCookie(w, &Id)
}

func (i *Id) get(r *http.Request) (*http.Cookie, error) {
	id, err := r.Cookie("id")
	if err == nil {
		return id, err
	} else {
		return nil, err
	}
}

func (i *Id) delete(w http.ResponseWriter) {
	Id := &http.Cookie{
		Name:     "id",
		Value:    "",
		MaxAge:   -1,
		Domain:   ls86.Server.Url,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	http.SetCookie(w, Id)
}

func (h *Hash) set(w http.ResponseWriter, value string) {

	hash := http.Cookie{
		Name:     "hash",
		Value:    value,
		Domain:   ls86.Server.Url,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3000,
	}
	http.SetCookie(w, &hash)
}

func (h *Hash) get(r *http.Request) (*http.Cookie, error) {
	hash, err := r.Cookie("hash")
	if err == nil {
		return hash, err
	} else {
		return nil, err
	}
}

func (h *Hash) delete(w http.ResponseWriter) {
	hash := &http.Cookie{
		Name:     "hash",
		Value:    "",
		MaxAge:   -1,
		Domain:   ls86.Server.Url,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	http.SetCookie(w, hash)
}

func (o *Other) CheckAuth(id, hash *http.Cookie) (*User, error) {

	u, err := ls86.Data.User.FindOne("id", id.Value)
	if err != nil {
		return nil, err
	} else {
		if u.AuthToken == hash.Value {
			return u, nil
		} else {
			return nil, ls86.Data.Error.TokenBad
		}
	}

}
