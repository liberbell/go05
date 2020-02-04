package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name        string `json: "Name"`
	PublicRepos int    `json: "public_repos"`
}

func userInfo(login string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user := &User{}
	dec := json.NewDecoder(resp.Body)
}
