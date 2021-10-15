package router

import (
	"todo-backend/app"

	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server *app.Server) {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

}
