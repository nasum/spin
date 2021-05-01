package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nasum/spin/lib"
	"github.com/nasum/spin/usecase"
)

type User struct {
	Name string `json:"name" xml:"name"`
}

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*lib.Context)
		twitterId, err := cc.GetTwitterID()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		if twitterId == "" {
			return c.JSON(http.StatusBadRequest, nil)
		}

		user := usecase.GetUserByTwitterAccountId(twitterId)

		u := &User{
			Name: user.Name,
		}

		return c.JSON(http.StatusOK, u)
	}
}
