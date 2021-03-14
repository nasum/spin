package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/nasum/spin/constants"
)

func getConfig() oauth1.Config {
	e := constants.GetEnv()
	return oauth1.Config{
		ConsumerKey:    e.TwitterAPIKey,
		ConsumerSecret: e.TwitterAPIKeySecret,
		CallbackURL:    e.CallbackURL,
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
}

func GetAuthorizationURL() (string, error) {
	config := getConfig()
	requestToken, _, err := config.RequestToken()

	if err != nil {
		return "", err
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)

	if err != nil {
		return "", err
	}

	return authorizationURL.String(), nil
}

func GetAccessToken(requestToken, verifier string) (string, string, error) {
	config := getConfig()
	_, requestTokenSecret, err := config.RequestToken()

	if err != nil {
		return "", "", err
	}

	accessToken, accessTokenSeacret, err := config.AccessToken(requestToken, requestTokenSecret, verifier)
	if err != nil {
		return "", "", err
	}

	return accessToken, accessTokenSeacret, nil
}

func GetTwitterClient(accessToken, accessTokenSeacret string) (*twitter.Client, error) {
	config := getConfig()
	httpClient := config.Client(oauth1.NoContext, oauth1.NewToken(accessToken, accessTokenSeacret))
	twitterClient := twitter.NewClient(httpClient)

	return twitterClient, nil
}

func GetAccount(client *twitter.Client) (*twitter.User, error) {
	accountVerifyParams := &twitter.AccountVerifyParams{
		IncludeEntities: twitter.Bool(false),
		SkipStatus:      twitter.Bool(true),
		IncludeEmail:    twitter.Bool(false),
	}

	account, _, err := client.Accounts.VerifyCredentials(accountVerifyParams)
	if err != nil {
		return nil, err
	}

	return account, nil
}
