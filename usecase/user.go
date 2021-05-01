package usecase

import (
	"github.com/nasum/spin/entity"
	"github.com/nasum/spin/repository"
)

func CreateUser(name, twitterAccountID, accessToken, accessTokenSeacret string) entity.User {
	user := repository.CreateUser(name, twitterAccountID, accessToken, accessTokenSeacret)
	return user
}

func DeleteUserByTwitterAccountId(twitterAccountId string) {
	repository.DeleteUser(twitterAccountId)
}

func GetUserByTwitterAccountId(twitterAccountId string) entity.User {
	user := repository.GetUser(twitterAccountId)
	return user
}
