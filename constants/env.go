package constants

import (
	"os"
)

type Env struct {
	TwitterAPIKey       string
	TwitterAPIKeySecret string
	CallbackURL         string
}

func GetEnv() Env {
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPIKeySecret := os.Getenv("TWITTER_API_KEY_SECRET")
	callbackURL := os.Getenv("CALLBACK_URL")

	e := Env{
		TwitterAPIKey:       twitterAPIKey,
		TwitterAPIKeySecret: twitterAPIKeySecret,
		CallbackURL:         callbackURL,
	}

	return e
}
