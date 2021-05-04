package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Redirect contains the required details for making a redirect.
type Redirect struct {
	Code int
	URL  string
}

// Mapping of redirects
var mappings = map[string]Redirect{
	//
	"www.mikemackintosh.com": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
	//
	"www.highonphp.com": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
	//
	"highonphp.com": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
}

// Start start package
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	e := echo.New()

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Pre(middleware.HTTPSNonWWWRedirect())

	// Listen on all requests
	e.Any("/*", func(c echo.Context) error {
		//validateDNS(c.Request().Host)
		// e.Pre(middleware.HTTPSNonWWWRedirect())
		redirect := mappings[c.Request().Host]
		c.Redirect(redirect.Code, redirect.URL)
		return nil
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
