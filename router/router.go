package router

import (
	"todo-backend/app"
	"todo-backend/controller"
	"todo-backend/helper"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server *app.Server) {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Validator = &helper.CustomValidator{Validator: validator.New()}
	todoController := controller.NewTodoController()

	todo := server.Echo.Group("/todo")
	todo.POST("/", todoController.Create)

}
