package router

import (
	"net/http"
	"todo-backend/app"
	"todo-backend/controller"
	"todo-backend/helper"
	"todo-backend/repository"
	"todo-backend/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server *app.Server) {
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Validator = &helper.CustomValidator{Validator: validator.New()}

	server.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Server is running")
	})

	todoRepository := repository.NewTodoRepository(server)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	todo := server.Echo.Group("/todos")
	todo.POST("/", todoController.Create)
	todo.GET("/", todoController.FindAll)

}
