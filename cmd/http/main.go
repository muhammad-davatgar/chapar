package main

import (
	messangerservice "chapar/internals/core/services/messanger"
	"chapar/internals/handlers/http"

	// "net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	msgService := messangerservice.NewMessangerService()

	HttpService := http.NewHttpService(msgService)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/ws", HttpService.ServeWs)

	app.Logger.Fatal(app.Start(":8080"))
}
