package main

import (
	"net/http"
	
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  
)

func main() {
// Echo Server
	e := echo.New()

// Middleware Setup
   e.Use(middleware.Recover())
	 e.Use(middleware.Logger())
// HTTPS Redirect
  e.Pre(middleware.HTTPSRedirect())
  
// Routing
 
 e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})
 
 
 e.Logger.Fatal(e.StartTLS(":443", "cert.pem", "key.pem"))
}
