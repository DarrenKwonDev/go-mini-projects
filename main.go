package main

import (
	"os"
	"strings"

	"github.com/DarrenKwonDev/learnGo/scrapper"
	"github.com/labstack/echo"
)

// Handler
func handleHome(c echo.Context) error {
	return c.File("home.html")
}

const fileName = "jobs.csv"

func handleScrape(c echo.Context) error {

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)

	defer os.Remove(fileName)
	// Attachment(뭘 줄 것인가, 무슨 이름으로 다운로드?)
	return c.Attachment(fileName, fileName)
}

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", handleHome)

	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
