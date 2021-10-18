package controller

import (
	"net/http"
	"todo-backend/model/web"
	"todo-backend/service"

	"github.com/labstack/echo/v4"
)

type TodoControllerImpl struct {
	Service service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &TodoControllerImpl{
		Service: service,
	}
}
func (controller *TodoControllerImpl) Create(c echo.Context) error {
	todo := new(web.TodoCreateRequest)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoResponse := controller.Service.Create(c.Request().Context(), *todo)
	return c.JSON(http.StatusOK, todoResponse)
}

func (controller *TodoControllerImpl) Update(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *TodoControllerImpl) Delete(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *TodoControllerImpl) FindByID(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *TodoControllerImpl) FindAll(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
