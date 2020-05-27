package main

import (
  "io"
  "html/template"
	"net/http"
  //"database/sql"
	
  //_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  
)

//Setting up echo templates
type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
// Load templates
  t := &Template{
    templates: template.Must(template.ParseGlob("html/*.html")),
  }

// Start echo Server
	e := echo.New()

// Render templates
  e.Renderer = t
 
// Middleware Setup
  e.Use(middleware.Recover())
  e.Use(middleware.Logger())

// HTTP to HTTPS Redirect
  e.Pre(middleware.HTTPSRedirect())
  
// Load statics at root
  e.Static("/", "html")
  
// Routing
  e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", template.HTML("<p>HTML Test</p>"))
	})

// Start server
  e.Logger.Fatal(e.StartTLS(":443", "cert.pem", "key.pem"))
}
