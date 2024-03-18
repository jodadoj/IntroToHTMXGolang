package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

type Count struct {
	Count int
}

type Contact struct {
		Name string
		Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name: name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("Clara", "cd@gmail.com"),
			newContact("John", "jd@gmail.com"),
		},
	}
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	
	// count := Count {Count: 0}
	data := newData()

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})
	
	// e.POST("/count", func(c echo.Context) error {
	// 	count.Count++
	// 	return c.Render(200, "count", count)
	// })

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		data.Contacts = append(data.Contacts, newContact(name, email))

		return c.Render(200, "display", data)
	})
	
	e.Logger.Fatal(e.Start(":42069"))
}