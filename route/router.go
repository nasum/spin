package route

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nasum/spin/handler"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t
	e.Static("/js", "dist/js")
	e.GET("/oauth/twitter", handler.SignUp())
	e.GET("/oauth/twitter/callback", handler.Callback())
	e.GET("*", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
