package main

import "fmt"

type User struct {
	Name        string `json: "Name"`
	PublicRepos int    `json: "public_repos"`
}

func userInfo(login string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)

}
