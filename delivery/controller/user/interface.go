package user

import "github.com/labstack/echo/v4"

type UserController interface {
	AddUser() echo.HandlerFunc
	ListAllUser() echo.HandlerFunc
	GetByID() echo.HandlerFunc
}
