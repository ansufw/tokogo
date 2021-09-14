package routes

import (
	"github.com/aysf/tokogo/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// not auth -> signup user, login user, get products,
	e.GET("/signup", controllers.CreateUserController)
	e.GET("/user", controllers.GetUserController)
	e.POST("/login", controllers.LoginUserController)

	// auth -> get user, get checkout, get transaction, post checkout/transaction
	e.GET("/user/:id", controllers.GetUserController)

	return e
}
