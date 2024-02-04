package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/spleeroosh/go-translate/handler"
)

type DB struct{}

func main() {
	app := echo.New()
	//var db DB
	// We could attach our db state to the handlers
	homeHandler := handler.HomeHandler{}
	translationsHandler := handler.TranslationsHandler{}

	app.Use(withUser)
	app.GET("/", homeHandler.HandleHomeShow)
	app.GET("/translations", translationsHandler.HandleTranslationsShow)

	app.Logger.Fatal(app.Start(":8080"))
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "rogovin.pavel@mail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
