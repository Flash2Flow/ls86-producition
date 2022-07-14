package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct{}

func (r Request) FindOneUserLogin(criterion string, login string, u interface{}) error {

	client := &http.Client{}
	url := fmt.Sprintf("%s%s%s/%s", rest.Url.BaseUrl, rest.Url.GetUser, criterion, login)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("rest-token", rest.Token)
	q := req.URL.Query()
	q.Add("criterion", criterion)
	q.Add("login", login)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	return json.NewDecoder(resp.Body).Decode(u)
}

func (r Request) FindOneUserEmail(criterion string, login string, u interface{}) error {

	client := &http.Client{}
	url := fmt.Sprintf("%s%s%s/%s", rest.Url.BaseUrl, rest.Url.GetUser, criterion, login)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("rest-token", rest.Token)
	q := req.URL.Query()
	q.Add("criterion", criterion)
	q.Add("login", login)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	return json.NewDecoder(resp.Body).Decode(u)
}
