package repository

import (
	"github.com/nasum/spin/entity"
	"github.com/nasum/spin/infrastructure"
	"github.com/nasum/spin/model"
)

func CreateUser(name, twitterAccountID, accessToken, accessTokenSeacret string) entity.User {
	conn := infrastructure.Connection

	user := model.User{
		Name:             name,
		TwitterAccountId: twitterAccountID,
	}
	twitterToken := model.TwitterToken{
		AccessToken:        accessToken,
		AccessTokenSeacret: accessTokenSeacret,
		User:               user,
	}

	conn.Create(&twitterToken)

	userEntity := entity.CreateUser(user.Name)
	return userEntity
}

func GetUser(twitterAccountId string) entity.User {
	conn := infrastructure.Connection
	var user model.User
	conn.First(&user, "twitter_account_id = ?", twitterAccountId)

	userEntity := entity.CreateUser(user.Name)
	return userEntity
}

func DeleteUser(twitterAccountId string) {
	conn := infrastructure.Connection
	var user model.User
	conn.First(&user, "twitter_account_id = ?", twitterAccountId)

	var twitterToken model.TwitterToken
	conn.First(&twitterToken, "user_id = ?", user.ID)

	conn.Delete(&twitterToken)
	conn.Delete(&user)
}
