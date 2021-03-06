package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/labstack/echo/v4"
	"github.com/nasum/spin/lib"
	"github.com/nasum/spin/twitter"
	"github.com/nasum/spin/usecase"
)

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
		cc := c.(*lib.Context)

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

		err = cc.StartSession(account.IDStr)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not start session")
		}

		user := usecase.GetUserByTwitterAccountId(account.IDStr)

		if user.Name == "" {
			usecase.CreateUser(account.ScreenName, account.IDStr, accessToken, accessTokenSeacret)
			http.Redirect(c.Response().Writer, c.Request(), "/home", http.StatusFound)
		} else {
			http.Redirect(c.Response().Writer, c.Request(), "/home", http.StatusFound)
		}

		return nil
	}
}
