package handler

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/labstack/echo/v4"
	"github.com/nasum/spin/constants"
)

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		e := constants.GetEnv()
		config := oauth1.Config{
			ConsumerKey:    e.TwitterAPIKey,
			ConsumerSecret: e.TwitterAPIKeySecret,
			CallbackURL:    e.CallbackURL,
			Endpoint:       twitter.AuthorizeEndpoint,
		}

		requestToken, _, err := config.RequestToken()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		authorizationURL, err := config.AuthorizationURL(requestToken)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter authorization url")
		}

		http.Redirect(c.Response().Writer, c.Request(), authorizationURL.String(), http.StatusFound)

		return nil
	}
}

func Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestToken, _, err := oauth1.ParseAuthorizationCallback(c.Request())

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not create twitter request token")
		}

		fmt.Println(requestToken)

		return nil
	}
}
