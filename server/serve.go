package server

import (
	"reidev-api/pages"

	"github.com/labstack/echo/v4"
)

func Serve() error {
	e := echo.New()

	e.GET("/", pages.HomePage)

	return e.Start(":1424")
}
