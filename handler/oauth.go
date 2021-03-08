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
		//conn := infrastructure.Connection

		e := constants.GetEnv()
		config := oauth1.Config{
			ConsumerKey:    e.TwitterAPIKey,
			ConsumerSecret: e.TwitterAPIKeySecret,
			CallbackURL:    e.CallbackURL,
			Endpoint:       twitterOAuth1.AuthorizeEndpoint,
		}

		requestToken, verifier, err := oauth1.ParseAuthorizationCallback(c.Request())

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		_, requestTokenSecret, err := config.RequestToken()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		accessToken, accessSecret, err := config.AccessToken(requestToken, requestTokenSecret, verifier)
		fmt.Println(accessToken)
		fmt.Println(accessSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(-1)
		}

		httpClient := config.Client(ctx, oauth1.NewToken(accessToken, accessSecret))
		twitterClient := twitter.NewClient(httpClient)

		accountVerifyParams := &twitter.AccountVerifyParams{
			IncludeEntities: twitter.Bool(false),
			SkipStatus:      twitter.Bool(true),
			IncludeEmail:    twitter.Bool(false),
		}

		user, _, err := twitterClient.Accounts.VerifyCredentials(accountVerifyParams)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(-1)
		}

		fmt.Println(user.ScreenName)

		return nil
	}
}
