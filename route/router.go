package route

import (
	"io"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nasum/spin/handler"
	"github.com/nasum/spin/infrastructure"
	"github.com/nasum/spin/lib"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() {
	t := &Template{
		templates: template.Must(template.ParseGlob("dist/*.html")),
	}

	infrastructure.Init()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &lib.Context{c}
			return next(cc)
		}
	})

	e.Renderer = t
	e.Static("/js", "dist/js")

	// OAuth
	oauth := e.Group("/oauth")
	oauth.GET("/twitter", handler.SignUp())
	oauth.GET("/twitter/callback", handler.Callback())

	// forward frontend router
	e.GET("*", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "")
	})
	e.Logger.Fatal(e.Start(":1323"))
	infrastructure.Close()
}
