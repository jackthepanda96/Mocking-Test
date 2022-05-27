package route

import (
	"go-clean/delivery/controller/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, cu user.UserController) {
	e.POST("/users", cu.AddUser())
	e.GET("/users", cu.ListAllUser())
	e.GET("/users/:id", cu.GetByID())
}
