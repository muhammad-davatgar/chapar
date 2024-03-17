package main

import (
	"chapar/internals/core/services/auth"
	messangerservice "chapar/internals/core/services/messanger"
	"chapar/internals/handlers/http"
	"chapar/internals/repositories/memuserdb"

	// "net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	msgService := messangerservice.NewMessangerService()

	memUserDB := memuserdb.NewMemoryUserDB()
	authService := auth.NewAuthenticationService(memUserDB)

	HttpService := http.NewHttpService(msgService, authService)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/ws", HttpService.ServeWs)
	app.POST("/signup", HttpService.SignUp)
	app.POST("/login", HttpService.Login)

	app.Logger.Fatal(app.Start(":8080"))
}
