package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/labstack/echo/v4"
	"github.com/nasum/spin/constants"
	"github.com/nasum/spin/infrastructure"
	"github.com/nasum/spin/model"
)

var Config *oauth1.Config

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		e := constants.GetEnv()
		Config := oauth1.Config{
			ConsumerKey:    e.TwitterAPIKey,
			ConsumerSecret: e.TwitterAPIKeySecret,
			CallbackURL:    e.CallbackURL,
			Endpoint:       twitterOAuth1.AuthorizeEndpoint,
		}

		requestToken, _, err := Config.RequestToken()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		authorizationURL, err := Config.AuthorizationURL(requestToken)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter authorization url")
		}

		http.Redirect(c.Response().Writer, c.Request(), authorizationURL.String(), http.StatusFound)

		return nil
	}
}

func Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		conn := infrastructure.Connection

		e := constants.GetEnv()
		config := oauth1.Config{
			ConsumerKey:    e.TwitterAPIKey,
			ConsumerSecret: e.TwitterAPIKeySecret,
			CallbackURL:    e.CallbackURL,
			Endpoint:       twitterOAuth1.AuthorizeEndpoint,
		}

		requestToken, verifier, err := oauth1.ParseAuthorizationCallback(c.Request())

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		_, requestTokenSecret, err := config.RequestToken()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		accessToken, accessSeacret, err := config.AccessToken(requestToken, requestTokenSecret, verifier)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter accessToken")
		}

		httpClient := config.Client(ctx, oauth1.NewToken(accessToken, accessSeacret))
		twitterClient := twitter.NewClient(httpClient)

		accountVerifyParams := &twitter.AccountVerifyParams{
			IncludeEntities: twitter.Bool(false),
			SkipStatus:      twitter.Bool(true),
			IncludeEmail:    twitter.Bool(false),
		}

		account, _, err := twitterClient.Accounts.VerifyCredentials(accountVerifyParams)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not get twitter user")
		}

		user := model.User{
			Name: account.ScreenName,
		}
		twitterToken := model.TwitterToken{
			AccessToken:        accessToken,
			AccessTokenSeacret: accessSeacret,
			User:               user,
		}

		conn.Create(&twitterToken)

		return nil
	}
}
