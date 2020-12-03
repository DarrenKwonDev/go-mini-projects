package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handler
func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", handleHome)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))

}

// scrapper.Scrape("python")
