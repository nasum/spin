package model

import (
	"gorm.io/gorm"
	"time"
)

type TwitterToken struct {
	Id                uint      `json:"id`
	UserId            uint      `json:"user_id`
	User              User      `json:"user`
	AccessToken       string    `json:"AccessToken"`
	AccessTokenSecret string    `json:"AccessTokenSecret"`
	created_at        time.Time `json:"created_at"`
	updated_at        time.Time `json:"updated_at"`
}
