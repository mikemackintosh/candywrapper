package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Redirect contains the required details for making a redirect.
type Redirect struct {
	Code int
	URL  string
}

// Start start package
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Pre(middleware.HTTPSNonWWWRedirect())

	// Listen on all requests
	e.Any("/*", func(c echo.Context) error {
		e.Logger.SetLevel(log.INFO)
		e.Logger.Info(map[string]string{"event": "request_received", "from_host": c.Request().Host})

		redirect := mappings["default"]
		if v, ok := mappings[c.Request().Host]; ok {
			redirect = v
		}

		c.Redirect(redirect.Code, redirect.URL)
		return nil
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
