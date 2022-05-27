package main

import (
	"go-clean/config"
	_userController "go-clean/delivery/controller/user"
	"go-clean/delivery/route"
	"go-clean/repository/user"
	_userService "go-clean/service/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	db := config.InitDB()
	config.Migrate(db)
	repo := user.New(db)
	srv := _userService.New(repo, validator.New())
	controll := _userController.New(srv)
	route.RegisterPath(e, controll)
	log.Fatal(e.Start(":8000"))
}
