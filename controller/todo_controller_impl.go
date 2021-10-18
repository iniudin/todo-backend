package controller

import "github.com/labstack/echo/v4"

type TodoControllerImpl struct {
}

func NewTodoController() TodoController {
	return &TodoControllerImpl{}
}
func (controller *TodoControllerImpl) Create(c echo.Context) error {
	panic("not implemented") // TODO: Implement
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
