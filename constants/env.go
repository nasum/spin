package constants

import (
	"os"
)

type Env struct {
	TwitterAPIKey       string
	TwitterAPIKeySecret string
	CallbackURL         string
	DataBaseURL         string
}

func GetEnv() Env {
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPIKeySecret := os.Getenv("TWITTER_API_KEY_SECRET")
	callbackURL := os.Getenv("CALLBACK_URL")
	databaseURL := os.Getenv("DATABASE_URL")

	e := Env{
		TwitterAPIKey:       twitterAPIKey,
		TwitterAPIKeySecret: twitterAPIKeySecret,
		CallbackURL:         callbackURL,
		DataBaseURL:         databaseURL,
	}

	return e
}
