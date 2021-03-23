package model

type User struct {
	Base
	Name             string `json:"name"`
	TwitterAccountId string `json:"TwitterAccountId"`
}
