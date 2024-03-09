package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/phatphumt/gohtmx/views/users"
)

type UserHandler struct{}

func (h UserHandler) HandleUser(c echo.Context) error {
	return render(c, users.Users())
}

func (h UserHandler) HandleSeeUser(c echo.Context) error {
	return render(c, users.SeeUsers())
}
