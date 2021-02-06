package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("dist/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/", "dist")
	e.GET("*", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
