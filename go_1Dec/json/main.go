package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User // avoid infinite recursion

	return json.Marshal(&struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  u.Name,
		Email: u.Email,
	})
}

func main() {
	u := User{Name: "Rahul", Email: "rahul@mail.com", Password: "secret"}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
