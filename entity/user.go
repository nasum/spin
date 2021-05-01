package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
}

func (u *User) ToJson() (string, error) {
	bytes, err := json.Marshal(u)

	if err != nil {
		return "", fmt.Errorf("%s: %v\n", os.Args[0], err)
	}

	return string(bytes), nil
}

func CreateUser(name string) User {
	var user User
	user.Name = name

	return user
}
