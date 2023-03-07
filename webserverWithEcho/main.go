package main

import (
	"scrapper"
	"strings"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.File("./home.html")
}

func handleScrape(c echo.Context) error {
	item := strings.ToLower(scrapper.CleanString(c.FormValue("samsung")))
	return c.File("home.html")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
