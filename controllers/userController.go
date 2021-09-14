package controllers

import (
	"net/http"
	"strconv"

	"github.com/aysf/tokogo/lib/database"
	"github.com/aysf/tokogo/models"
	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	userInput := new(models.User)
	if err := c.Bind(userInput); err != nil {
		return err
	}
	user, err := database.CreateUser(userInput)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success created user",
		"data":    user,
	})
}

func GetUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := database.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user data",
		"data":    user,
	})
}
