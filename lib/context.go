package lib

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func (c *Context) StartSession(twitterID string) error {
	sess, err := session.Get("spin_session", c)

	if err != nil {
		return fmt.Errorf("%s: %v\n", os.Args[0], err)
	}

	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	}

	sess.Values["TwitterID"] = twitterID

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("%s: %v\n", os.Args[0], err)
	}

	return nil
}

func (c *Context) ClearSession() error {
	sess, err := session.Get("spin_session", c)

	if err != nil {
		return fmt.Errorf("%s: %v\n", os.Args[0], err)
	}

	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}

	return nil
}

func (c *Context) GetTwitterID() (string, error) {
	sess, err := session.Get("spin_session", c)

	if err != nil {
		return "", fmt.Errorf("%s: %v\n", os.Args[0], err)
	}

	if sess.Values["TwitterID"] == nil {
		return "", nil
	}

	return sess.Values["TwitterID"].(string), nil
}
