package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: can't call api")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	user := User{Login: login}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	user, err := userInfo("dheerapat")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
