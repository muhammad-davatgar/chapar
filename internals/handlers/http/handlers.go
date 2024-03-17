package http

import (
	"chapar/internals/core/domain"
	"log"

	"github.com/labstack/echo/v4"
)

func (app *HTTPServer) SignUp(c echo.Context) error {
	entryCredentials := domain.EntryCredentials{}

	err := c.Bind(&entryCredentials)
	if err != nil {
		return echo.ErrBadRequest
	}

	token, err := app.authentication.SignUp(entryCredentials)
	if err != nil {
		// TODO : error handling
		log.Println(err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(201, map[string]any{"token": token})
}

func (app *HTTPServer) Login(c echo.Context) error {
	entryCredentials := domain.EntryCredentials{}

	err := c.Bind(&entryCredentials)
	if err != nil {
		return echo.ErrBadRequest
	}

	token, err := app.authentication.Login(entryCredentials)
	if err != nil {
		// TODO : error handling
		return echo.ErrInternalServerError
	}

	return c.JSON(201, map[string]any{"token": token})
}
