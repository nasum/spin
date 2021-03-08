package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         uint      `json:"id`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}
