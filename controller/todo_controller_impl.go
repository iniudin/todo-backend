package controller

import (
	"net/http"
	"strconv"
	"todo-backend/helper"
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
	todo := new(web.TodoUpdateRequest)
	todoID, err := strconv.Atoi(c.Param("todoID"))
	helper.PanicIfError(err)
	todo.ID = uint(todoID)

	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoResponse := controller.Service.Update(c.Request().Context(), *todo)
	return c.JSON(http.StatusOK, todoResponse)
}

func (controller *TodoControllerImpl) Delete(c echo.Context) error {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	helper.PanicIfError(err)

	controller.Service.Delete(c.Request().Context(), uint(todoID))
	helper.PanicIfError(err)

	return c.JSON(http.StatusOK, "Todo deleted")
}

func (controller *TodoControllerImpl) FindByID(c echo.Context) error {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	helper.PanicIfError(err)

	todoResponse := controller.Service.FindByID(c.Request().Context(), uint(todoID))
	return c.JSON(http.StatusOK, todoResponse)
}

func (controller *TodoControllerImpl) FindAll(c echo.Context) error {
	todoResponses := controller.Service.FindAll(c.Request().Context())
	return c.JSON(http.StatusOK, todoResponses)
}
