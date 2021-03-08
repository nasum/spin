package model

type TwitterToken struct {
	Base
	UserId             int    `json:"UserId"`
	User               User   `gorm:"foreignKey:UserId"`
	AccessToken        string `json:"AccessToken"`
	AccessTokenSeacret string `json:"AccessTokenSeacret"`
}
