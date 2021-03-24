package usecase

import (
	"os"
	"testing"

	"github.com/nasum/spin/infrastructure"
)

const TWITTER_ACCOUNT_ID = "1234"

func TestMain(m *testing.M) {
	infrastructure.Init()
	CreateUser("test", TWITTER_ACCOUNT_ID, "accessToken", "accessTokenSeacret")
	exitCode := m.Run()
	DeleteUserByTwitterAccountId(TWITTER_ACCOUNT_ID)
	infrastructure.Close()
	os.Exit(exitCode)
}

func TestGetUserByTwitterAccountId(t *testing.T) {
	user := GetUserByTwitterAccountId(TWITTER_ACCOUNT_ID)

	if user.TwitterAccountId != TWITTER_ACCOUNT_ID {
		t.Fatal("Can not create user: actuary")
	}
}
