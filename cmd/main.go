package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phatphumt/gohtmx/handlers"
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api")

	userHandler := handlers.UserHandler{}
	e.GET("/", userHandler.HandleUser)
	e.GET("/see", userHandler.HandleSeeUser)
	e.Logger.Fatal(e.Start(":1323"))
}
