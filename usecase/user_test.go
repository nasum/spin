package usecase

import (
	"os"
	"testing"

	"github.com/nasum/spin/infrastructure"
	"github.com/stretchr/testify/assert"
)

const TWITTER_ACCOUNT_ID = "1234"

func TestMain(m *testing.M) {
	infrastructure.Init()
	_ = CreateUser("test", TWITTER_ACCOUNT_ID, "accessToken", "accessTokenSeacret")
	exitCode := m.Run()
	DeleteUserByTwitterAccountId(TWITTER_ACCOUNT_ID)
	infrastructure.Close()
	os.Exit(exitCode)
}

func TestGetUserByTwitterAccountId(t *testing.T) {
	user := GetUserByTwitterAccountId(TWITTER_ACCOUNT_ID)

	if user.Name != "test" {
		t.Fatal("Can not create user: actuary")
	}

	assert.Equal(t, user.Name, "test")
}
