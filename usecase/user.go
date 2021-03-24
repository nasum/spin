package usecase

import (
	"github.com/nasum/spin/infrastructure"
	"github.com/nasum/spin/model"
)

func CreateUser(name, twitterAccountID, accessToken, accessTokenSeacret string) {
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
}

func DeleteUserByTwitterAccountId(twitterAccountId string) {
	conn := infrastructure.Connection
	user := GetUserByTwitterAccountId(twitterAccountId)
	twitterToken := GetTwitterTokenByUserId(user.ID)
	conn.Delete(&twitterToken)
	conn.Delete(&user)
}

func GetUserByTwitterAccountId(twitterAccountId string) model.User {
	conn := infrastructure.Connection
	var user model.User
	conn.First(&user, "twitter_account_id = ?", twitterAccountId)
	return user
}

func GetTwitterTokenByUserId(userId uint) model.TwitterToken {
	conn := infrastructure.Connection
	var twitterToken model.TwitterToken
	conn.First(&twitterToken, "user_id = ?", userId)
	return twitterToken
}
