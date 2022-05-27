package user

import (
	"go-clean/entities"
	"go-clean/service/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllerUser struct {
	service user.UserService
}

func New(srv user.UserService) *controllerUser {
	return &controllerUser{
		service: srv,
	}
}

func (cu controllerUser) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data entities.RequestUser

		err := c.Bind(&data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"message": "error input format",
				"data":    nil,
			})
		}

		res, err := cu.service.Add(data)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"message": "something wrong with your data",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"success": true,
			"message": "user created",
			"data":    res,
		})
	}
}

func (cu controllerUser) ListAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cu.service.ListUsers()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"message": "something wrong with your server",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "success get data",
			"data":    res,
		})
	}
}

func (cu controllerUser) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param(":id")

		convID, _ := strconv.Atoi(id)

		res, err := cu.service.MyProfile(convID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"message": "something wrong with your server",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "success get data",
			"data":    res,
		})
	}
}
