package routes

import (
	"github.com/aysf/tokogo/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/signup", controllers.CreateUserController)
	e.GET("/user", controllers.GetUserController)

	return e
}
