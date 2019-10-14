package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Message struct
type Message struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	m := &Message{
		Title:   "Greetings",
		Message: "Hello World!",
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, m)
	})
	e.Logger.Fatal(e.Start(":80"))
}
