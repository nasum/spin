package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/labstack/echo/v4"
	"github.com/nasum/spin/infrastructure"
	"github.com/nasum/spin/model"
	"github.com/nasum/spin/twitter"
)

var Config *oauth1.Config

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationURL, err := twitter.GetAuthorizationURL()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter authorization url")
		}

		http.Redirect(c.Response().Writer, c.Request(), authorizationURL, http.StatusFound)

		return nil
	}
}

func Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := infrastructure.Connection

		requestToken, verifier, err := oauth1.ParseAuthorizationCallback(c.Request())

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		accessToken, accessTokenSeacret, err := twitter.GetAccessToken(requestToken, verifier)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not get access token")
		}

		client, err := twitter.GetTwitterClient(accessToken, accessTokenSeacret)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter client")
		}

		account, err := twitter.GetAccount(client)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Cant not get account")
		}

		user := model.User{
			Name: account.ScreenName,
		}
		twitterToken := model.TwitterToken{
			AccessToken:        accessToken,
			AccessTokenSeacret: accessTokenSeacret,
			User:               user,
		}

		conn.Create(&twitterToken)

		return nil
	}
}
